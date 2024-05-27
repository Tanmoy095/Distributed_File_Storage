package p2p

import (
	"fmt"
	"net"
	"sync"
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

type TCPTransportOpts struct {
	ListenAddress string
	ShakeHandFunc HandshakeFunc
	Decoder       Decoder
}
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	mu       sync.Mutex
	peers    map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}

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

type Temp struct {
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.ShakeHandFunc(peer); err != nil {
		conn.Close()
		fmt.Print("TCP handshake error: %s\n", err)
		return

	}
	//Read Loop
	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue

		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("message:%+v\n", msg)
	}

}
