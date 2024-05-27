package p2p

// peer is a interface that represent the remote node
type peer interface{}

// transport is anything that handles the communication
// between the nodes in the network ..this can be of the from
// (tcp,websocket,anything)
type transport interface{}
