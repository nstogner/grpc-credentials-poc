package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/nstogner/grpc-auth-poc/poc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:50051"))
	if err != nil {
		log.Fatalf("[server] failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("[server] failed to load credentials: %v", err)
	}

	svr := grpc.NewServer(grpc.Creds(creds))
	poc.RegisterPocServer(svr, &server{})

	log.Println("[server] starting...")
	log.Fatal(svr.Serve(lis))
}

type server struct{}

func (s *server) Hey(ctx context.Context, req *poc.Req) (*poc.Rep, error) {
	log.Println("[server] got a request!")
	return &poc.Rep{}, nil
}
