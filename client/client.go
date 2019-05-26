package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Perform the address resolution and also
	// specialize the socket to only be able
	// to read and write to and from such
	// resolved address.
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("Unable to create connection to server: %v", err)
	}
	defer conn.Close()

	n, err := fmt.Fprintf(conn, "something")
	if err != nil {
		log.Fatalf("Unable to write %v to server: %v", n, err)
	}
	fmt.Printf("Wrote %v to server\n", n)
}
