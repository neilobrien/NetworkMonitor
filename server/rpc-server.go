package main

import (
	"context"
	"fmt"
	"log"
	"net"

	ppb "github.com/neilobrien/NetworkMonitor/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *ppb.HelloRequest) (*ppb.HelloResponse, error) {
	fmt.Printf("Hello function was invoked with %v\n", req)
	firstName := req.GetHello().GetFirstName()
	result := "Hello " + firstName
	res := &ppb.HelloResponse{
		Result: result,
	}
	return res, nil
}

const serverPort = ":50051"
const serverIP = "0.0.0.0"
const serverProtocol = "tcp"

func main() {

	lis, err := net.Listen(serverProtocol, serverIP+serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Listening on %v%v...\n", serverIP, serverPort)

	opts := []grpc.ServerOption{}
	tls := false
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	ppb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
