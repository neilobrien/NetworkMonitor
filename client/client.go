package main

import (
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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

	// n, err := fmt.Fprintf(conn, "something")
	// if err != nil {
	// 	log.Fatalf("Unable to write %v to server: %v", n, err)
	// }
	// fmt.Printf("Wrote %v to server\n", n)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	gopacket.SerializeLayers(buf, opts,
		&layers.Ethernet{},
		&layers.IPv4{
			SrcIP: net.IP{1, 2, 3, 4},
			DstIP: net.IP{5, 6, 7, 8}},
		&layers.UDP{
			DstPort: layers.UDPPort(8888),
		})
	w, err := conn.Write(buf.Bytes())
	if err != nil {
		log.Fatalf("Unable to write to server: %v", err)
	}
	log.Printf("Sent %v to server", w)
}
