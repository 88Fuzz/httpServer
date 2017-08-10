package httpServer

import "net"
import "io"
import "time"
import "bytes"

const timeout = 300 * time.Millisecond

type SingleConnectionProcessor struct{}

func (connectionProcessor SingleConnectionProcessor) Init() error {
	//Do nothing
	return nil
}

func (connectionProcessor SingleConnectionProcessor) Finish() {
	//Do nothing
}

func (connectionProcessor SingleConnectionProcessor) Process(connection net.Conn,
	processorProvider HttpRequestProcessorProvider) {
	request, err := readRequest(connection)
	if err != nil {
		writeError(connection, BAD_REQUEST)
		return
	}

	response := processRequest(processorProvider, request)
	responseString := createHttpResponse(request, response)
	writeAndClose(connection, responseString)
}

func readRequest(connection net.Conn) (Request, error) {
	var emptyRequest Request
	var requestBuffer bytes.Buffer
	var err error
	count, readSize := 256, 256

	for count >= readSize {
		bytes := make([]byte, readSize)
		err = connection.SetReadDeadline(time.Now().Add(timeout))
		count, err = connection.Read(bytes)
		_, err = requestBuffer.Write(bytes)

		if err != nil {
			break
		}
	}

	if err != nil {
		if netError, ok := err.(net.Error); ok {
			if !netError.Timeout() {
				return emptyRequest, err
			}
		} else if err != io.EOF {
			return emptyRequest, err
		}
	}

	reqStr := requestBuffer.String()
	return parseRequest(reqStr)
}

func writeError(connection net.Conn, statusCode StatusCode_t) {
	var errorResponse Response
	var request Request
	errorResponse.StatusCode = statusCode
	request.RequestType = FULL

	responseString := createHttpResponse(request, errorResponse)
	writeAndClose(connection, responseString)
}

func writeAndClose(connection net.Conn, response string) {

	err := connection.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		//swallow error and close the connection
		connection.Close()
		return
	}
	connection.Write([]byte(response))
	connection.Close()
}
