package httpServer

import "errors"
import "net"

type inboundConnection struct {
	connection            net.Conn
	connectionProcessor   ConnectionProcessor
	httpProcessorProvider HttpRequestProcessorProvider
}

type ThreadedConnectionProcessor struct {
	NumberOfThreads int
	Processor       ConnectionProcessor
	channel         chan<- inboundConnection
}

func (connectionProcessor *ThreadedConnectionProcessor) Init() error {
	if connectionProcessor.NumberOfThreads < 1 {
		return errors.New("NumberOfThreads cannot be less than 1.")
	} else if connectionProcessor.Processor == nil {
		return errors.New("Processor cannot be null.")
	}

	channel := make(chan inboundConnection, 100)
	for i := 0; i < connectionProcessor.NumberOfThreads; i++ {
		go processChannel(channel)
	}
	connectionProcessor.channel = channel

	return nil
}

func processChannel(channel <-chan inboundConnection) {
	for inConnection := range channel {
		inConnection.connectionProcessor.Process(
			inConnection.connection, inConnection.httpProcessorProvider)
	}
}

func (connectionProcessor *ThreadedConnectionProcessor) Finish() {
	close(connectionProcessor.channel)
}

func (connectionProcessor *ThreadedConnectionProcessor) Process(connection net.Conn,
	processorProvider HttpRequestProcessorProvider) {
	var data inboundConnection
	data.connection = connection
	data.connectionProcessor = connectionProcessor.Processor
	data.httpProcessorProvider = processorProvider
	connectionProcessor.channel <- data
}
