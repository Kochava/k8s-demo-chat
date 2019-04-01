package tcputil

import "net"

// Handler defines a common connection handler
type Handler interface {
	Handle(net.Conn)
}
