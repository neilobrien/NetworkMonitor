package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	ppb "github.com/neilobrien/NetworkMonitor/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

const (
	serverIP       = "0.0.0.0"
	serverPort     = ":50051"
	serverProtocol = "tcp"
	fpath          = "../probes"
	fname          = "probe.txt"
)

func (*server) Hello(ctx context.Context, req *ppb.HelloRequest) (*ppb.HelloResponse, error) {
	fmt.Printf("Hello function was invoked with %v\n", req)
	firstName := req.GetHello().GetFirstName()
	result := "Hello " + firstName
	res := &ppb.HelloResponse{
		Result: result,
	}
	return res, nil
}

var p = &ppb.CommonProbes{}

func (*server) GetProbesVersion(ctx context.Context, req *ppb.VersionRequest) (*ppb.CommonProbes, error) {
	fmt.Printf("Returning Version\n")
	res := &ppb.CommonProbes{
		Version: p.GetVersion(),
	}
	return res, nil
}

func fromFile(fname string, pb proto.Message) {
	i := 0
	for {
		i++
		in, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatalf("Can't read file: %v", err)
		}
		fmt.Println("debugging...")
		fmt.Println(string(in))

		err = proto.UnmarshalText(string(in), pb)
		if err != nil {
			log.Fatalf("Can't read from file: %v", err)
		}
		fmt.Printf("Sleep...%d\n", i)
		time.Sleep(5 * time.Second)
		fmt.Println("Awake...")
	}
}

func main() {
	// p := &ppb.CommonProbes{}
	go fromFile(strings.Join([]string{fpath, fname}, "/"), p)

	fmt.Printf("Version: %v\n", p.GetVersion())
	fmt.Println("reading probes....")

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
