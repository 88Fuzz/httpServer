package main

import "os"
import "fmt"
import "io/ioutil"
import "github.com/88Fuzz/httpServer"

func main() {
	listener, err := httpServer.StartServer(8080)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error could not start up: %s", err.Error())
		return
	}

	httpServer.Process(listener, httpServer.StdoutConnectionProcessor{}, FuckOffProvider{})
}

type FuckOffProcessor struct{}

func (fuckOff FuckOffProcessor) Get(request httpServer.Request) httpServer.Response {
	var response httpServer.Response
	byteArray, err := ioutil.ReadFile("index.html")
	if err != nil {
		response.StatusCode = httpServer.INTERNAL_SERVER_ERROR
		return response
	}

	response.Body = string(byteArray)
	response.StatusCode = httpServer.OK

	return response
}

func (fuckOff FuckOffProcessor) Head(request httpServer.Request) httpServer.Response {
	var response httpServer.Response
	response.StatusCode = httpServer.NOT_IMPLEMENTED

	return response
}

func (fuckOff FuckOffProcessor) Post(request httpServer.Request) httpServer.Response {
	var response httpServer.Response
	response.StatusCode = httpServer.NOT_IMPLEMENTED

	return response
}

func (fuckOff FuckOffProcessor) ExtensionMethod(request httpServer.Request) httpServer.Response {
	var response httpServer.Response
	response.StatusCode = httpServer.NOT_IMPLEMENTED

	return response
}

type FuckOffProvider struct{}

func (fuckOff FuckOffProvider) GetHttpRequestProcessor() (httpServer.HttpRequestProcessor, error) {
	return FuckOffProcessor{}, nil
}

func (fuckOff FuckOffProvider) HttpRequestProcessorDone(requestProcessor httpServer.HttpRequestProcessor) {
}
