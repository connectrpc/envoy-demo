// Copyright 2021-2023 The Connect Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//go:generate buf generate

import (
	"context"
	"flag"
	"log"
	"net"

	pingv1 "github.com/connectrpc/envoy-demo/internal/gen/ping/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	_ "github.com/connectrpc/envoy-demo/internal/codec/json"
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
