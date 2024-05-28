package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	Opts := TCPTransportOpts{
		ListenAddress: ":3000",
		ShakeHandFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(Opts)

	assert.Equal(t, tr.ListenAddress, ":3000")
	//server
	assert.Nil(t, tr.ListenAndAccept())

}
