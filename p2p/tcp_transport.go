package p2p

import (
	"fmt"
	"net"
)

// tcp peer represent a remote node over a tcp established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	connection net.Conn
	// if we dial and retrive a connection =>outbound == true
	// if we recive and retrive a connection =>outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		connection: conn,
		outbound:   outbound,
	}
}

// close implements the peer interface
func (p TCPPeer) Close() error {
	return p.connection.Close()

}

type TCPTransportOpts struct {
	ListenAddress string
	ShakeHandFunc HandshakeFunc
	Decoder       Decoder
	onpeer        func(Peer) error
}
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcchn   chan RPC
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcchn:           make(chan RPC),
	}

}

// we can only read  from the channel
// consume implements the Transport interface
// for reading incoming messages received from another peer in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcchn
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err

	}
	go t.StartAcceptLoop()
	return nil

}

func (t *TCPTransport) StartAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error : %s\n", err)

		}
		fmt.Printf("new incoming connection %v\n", conn)

		go t.HandleConnection(conn)
	}
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	var err error
	defer func() {
		fmt.Printf("dropping peer connection %s", err)
		conn.Close()

	}()

	// basically here what happens is do an handshake if that ok we gonna do the onpeer if that fails
	//we gonna drop . if that continue ..than start our read lopp
	peer := NewTCPPeer(conn, true)

	if err = t.ShakeHandFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return

	}
	//if somebody provides onpeer function
	if t.onpeer != nil {
		if err = t.onpeer(peer); err != nil {
			return

		}

	}
	//Read Loop
	rpc := RPC{}
	for {
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue

		}
		rpc.From = conn.RemoteAddr()
		t.rpcchn <- rpc
	}

}
