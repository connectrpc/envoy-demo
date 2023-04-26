package main

//go:generate buf generate

import (
	"context"
	"flag"
	"log"
	"net"

	pingv1 "github.com/bufbuild/connect-envoy-demo/internal/gen/ping/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	_ "github.com/bufbuild/connect-envoy-demo/internal/codec/json"
	_ "google.golang.org/grpc/encoding/gzip"
)

type PingServer struct {
	pingv1.UnimplementedPingServiceServer
}

func (ps *PingServer) Ping(
	ctx context.Context,
	req *pingv1.PingRequest,
) (*pingv1.PingResponse, error) {
	// You can set headers like Cache-Control.
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"Cache-Control": "max-age=604800",
	}))

	return &pingv1.PingResponse{
		Number: req.Number,
	}, nil
}

func main() {
	listenAddress := flag.String("listen", ":8080", "Address to listen on")
	flag.Parse()

	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("Error listening on %q: %v", *listenAddress, err)
	}
	log.Printf("Listening on %v", listener.Addr())

	server := grpc.NewServer()
	pingv1.RegisterPingServiceServer(server, &PingServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error serving connections: %v", err)
	}
}
