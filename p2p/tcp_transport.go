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

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	mu            sync.Mutex
	peers         map[net.Addr]peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
		shakeHands:    NOPHandshakeFunc,
	}

}
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
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
		go t.HandleConnection(conn)
	}
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	fmt.Printf("new incoming connection %v\n", peer)
}
