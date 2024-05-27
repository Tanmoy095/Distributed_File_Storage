package main

import (
	"log"

	"github.com/filestore/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		ShakeHandFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)

	}
	select {}
}
