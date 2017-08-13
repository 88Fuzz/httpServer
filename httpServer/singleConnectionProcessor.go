package httpServer

import "net"
import "io"
import "time"
import "bytes"

const keepAliveTimeout = 5000 * time.Millisecond
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
	keepProcessing := true
	keepAliveCloseTime := time.Now().Add(keepAliveTimeout)
	for keepProcessing {
		if err := connection.SetDeadline(keepAliveCloseTime); err != nil {
			connection.Close()
			return
		}

		request, err, closeConnection := readRequest(connection, keepAliveCloseTime)
		if closeConnection {
			//No information asked by client. Close and do nothing.
			connection.Close()
			return
		}
		if err != nil {
			writeError(connection, BAD_REQUEST)
			return
		}

		response := processRequest(processorProvider, request)
		responseString := createHttpResponse(request, response)
		keepProcessing = writeAndClose(connection, responseString, false)
	}
}

func readRequest(connection net.Conn, maxTimeout time.Time) (Request, error, bool) {
	var emptyRequest Request
	var requestBuffer bytes.Buffer
	var err error
	count, readSize := 256, 256

	for count >= readSize {
		bytes := make([]byte, readSize)
		err = connection.SetReadDeadline(time.Now().Add(timeout))
		count, err = connection.Read(bytes)
		if count != 0 {
			_, err = requestBuffer.Write(bytes)
		} else if err == io.EOF {
			//No bytes were read and nothing was read. Keep waiting for a valid read
			count = readSize
			err = nil
		}

		if time.Now().After(maxTimeout) {
			if requestBuffer.Len() > 0 {
				break
			} else {
				return emptyRequest, nil, true
			}
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		if netError, ok := err.(net.Error); ok {
			if !netError.Timeout() {
				return emptyRequest, err, false
			}
		} else if err != io.EOF {
			return emptyRequest, err, false
		}
	}

	reqStr := requestBuffer.String()
	request, err := parseRequest(reqStr)
	return request, err, false
}

func writeError(connection net.Conn, statusCode StatusCode_t) {
	var errorResponse Response
	var request Request
	errorResponse.StatusCode = statusCode
	request.RequestType = FULL

	responseString := createHttpResponse(request, errorResponse)
	writeAndClose(connection, responseString, true)
}

func writeAndClose(connection net.Conn, response string, closeConnection bool) bool {
	err := connection.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		//swallow error and close the connection
		connection.Close()
		return false
	}
	connection.Write([]byte(response))
	//connection.Write([]byte("\000"))
	//connection.Close()
	if closeConnection {
		connection.Close()
		return false
	}

	one := []byte{}
	connection.SetReadDeadline(time.Now())
	if _, err = connection.Read(one); err != nil {
		if err == io.EOF {
			connection.Close()
			return false
		}
		if err, ok := err.(net.Error); ok && err.Timeout() {
			//Error was timeout, which is to be expected if connection is open
			return true
		}
	}
	return false
}
