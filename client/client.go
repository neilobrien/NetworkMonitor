package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/golang/protobuf/proto"
	ppb "github.com/neilobrien/NetworkMonitor/protos"

	"google.golang.org/grpc"
)

var (
	serverIP = "35.246.29.114"
	p        = &ppb.CommonProbes{}
)

const (
	serverPort     = ":50051"
	serverProtocol = "tcp"
	fpath          = "probes"
	fname          = "client_probe.txt"
)

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
	getVersion(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)

	// doUnaryWithDeadline(c, 5*time.Second) // should complete
	// doUnaryWithDeadline(c, 1*time.Second) // should timeout
	go fromFile(strings.Join([]string{fpath, fname}, "/"), p)
	time.Sleep(60 * time.Second)
}

func getVersion(c ppb.HelloServiceClient) {
	fmt.Println("Requesting Version...")
	req := &ppb.VersionRequest{}
	res, err := c.GetProbesVersion(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GetProbesVersion RPC: %v", err)
	}
	log.Printf("Received Version %v", res.Version)
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

func fromFile(fname string, pb proto.Message) {
	fmt.Println("Reading client probes...")
	i := 0
	for {
		i++
		in, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatalf("Can't read file: %v", err)
		}
		// fmt.Println("debugging...")
		// fmt.Println(string(in))

		err = proto.UnmarshalText(string(in), pb)
		if err != nil {
			log.Fatalf("Can't read from file: %v", err)
		}
		fmt.Printf("Client Version: %v\n", p.GetVersion())
		fmt.Printf("Sleep...%d\n", i)
		time.Sleep(5 * time.Second)
		fmt.Println("Awake...")
	}
}
