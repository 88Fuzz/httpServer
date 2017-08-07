package httpServer

import "fmt"
import "net"
import "errors"
import "os"

func StartServer(port int) (net.Listener, error) {
	if port < 0 || port > 65535 {
		return nil, errors.New("Port is not in valid range [0, 65535]")
	}

	portStr := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", portStr)

	if err != nil {
		return nil, err
	}

	return listener, err
}

func Process(listener net.Listener, connectionProcessor ConnectionProcessor,
	requestProcessorProvider HttpRequestProcessorProvider) {
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection.\n")
			continue
		}

		connectionProcessor.Process(connection, requestProcessorProvider)
	}
}
