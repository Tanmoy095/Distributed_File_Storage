package p2p

// handshake func...
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
