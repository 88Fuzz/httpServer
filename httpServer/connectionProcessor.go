package httpServer

import "net"

type ConnectionProcessor interface {
	Process(net.Conn) error
}
