FROM docker.io/golang:alpine as builder
WORKDIR /workspace
COPY go.mod go.sum server.go /workspace/
COPY internal /workspace/internal
RUN CGO_ENABLED=0 go build -o /go/bin/server .

FROM docker.io/alpine
COPY --from=builder /go/bin/server /usr/local/bin/server
ENTRYPOINT ["/usr/local/bin/server"]
