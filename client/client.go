package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc/credentials"

	ppb "github.com/neilobrien/NetworkMonitor/protos"

	"google.golang.org/grpc"
)

var serverIP = "35.246.29.114"
var serverPort = ":50051"
var serverProtocol = "tcp"

func main() {
	gRPCTarget := flag.Bool("local", false, "gRPC Target")
	flag.Parse()
	if *gRPCTarget {
		serverIP = "0.0.0.0"
	}

	fmt.Println(*gRPCTarget)
	fmt.Println(serverIP)
	fmt.Println("Hello I'm a client")

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial(serverIP+serverPort, opts)
	fmt.Printf("Attempting to connect to %s%s...\n\n", serverIP, serverPort)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := ppb.NewHelloServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)

	// doUnaryWithDeadline(c, 5*time.Second) // should complete
	// doUnaryWithDeadline(c, 1*time.Second) // should timeout
}

func doUnary(c ppb.HelloServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &ppb.HelloRequest{
		Hello: &ppb.Hello{
			FirstName: "Neil",
			LastName:  "O'Brien",
		},
	}
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
