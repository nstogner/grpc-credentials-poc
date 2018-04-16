package main

import (
	"context"
	"log"

	"github.com/nstogner/grpc-credentials-poc/poc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		log.Fatalf("[client] unable to load credentials: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("[client] unable to dial server: %v", err)
	}
	defer conn.Close()

	client := poc.NewPocClient(conn)

	log.Println("[client] requesting server")
	if _, err = client.Hey(context.Background(), &poc.Req{}); err != nil {
		log.Fatalf("[client] unable to execute request: %v", err)
	}

	log.Println("[client] done")
}
