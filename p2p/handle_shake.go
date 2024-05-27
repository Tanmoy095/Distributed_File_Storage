package p2p

// handshake func...
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
