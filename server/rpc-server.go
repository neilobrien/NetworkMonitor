package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"github.com/neilobrien/NetworkMonitor/greet"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greet.GreetResponse{
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
	greet.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
