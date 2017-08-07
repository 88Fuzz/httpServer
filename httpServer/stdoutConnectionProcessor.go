package httpServer

import "net"
import "fmt"
import "io"

import "time"
import "bytes"

type StdoutConnectionProcessor struct{}

func (connectionProcessor StdoutConnectionProcessor) Process(connection net.Conn) error {
	var requestBuffer bytes.Buffer
	var err error
	count, readSize := 256, 256
	timeout := 50 * time.Millisecond

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
				return err
			}
		} else if err != io.EOF {
			return err
		}
	}

	reqStr := requestBuffer.String()
	request, err := parseRequest(reqStr)

	if err != nil {
		fmt.Printf("Error parsing ")
		fmt.Println(err)
	} else {
		fmt.Println("Valid request: ", request.requestType, request.version, request.method, request.methodValue, request.path,
			request.headers, request.body)
	}

	fmt.Printf("All done here\n")

	returnString := createHttpResponse(request, createTempResponse())
	fmt.Println("Writing ", returnString)
	connection.Write([]byte(returnString))
	connection.Close()

	return nil
}

func createTempResponse() Response {
	var response Response
	response.statusCode = NOT_FOUND

	return response
}
