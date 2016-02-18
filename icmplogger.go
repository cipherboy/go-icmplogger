package main

import "golang.org/x/net/icmp"

import (
	"log"
)

func main() {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		rb := make([]byte, 1500)
		_, peer, err := c.ReadFrom(rb)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("received ping from %v", peer)
	}
}
