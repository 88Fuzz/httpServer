package httpServer

import "net"

type ConnectionProcessor interface {
	Init() error
	Finish()
	Process(net.Conn, HttpRequestProcessorProvider)
}
