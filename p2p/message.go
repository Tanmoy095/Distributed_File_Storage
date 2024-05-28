package p2p

import "net"

//message holds any arbitrary data that is being send over each transport between two nodes in the network

type RPC struct {
	From    net.Addr
	Payload []byte
}
