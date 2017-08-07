package httpServer

type HttpRequestProcessor interface {
	Get(req Request) Response
	Head(req Request) Response
	Post(req Request) Response
	ExtensionMethod(req Request) Response
}

type HttpRequestProcessorProvider interface {
	GetHttpRequestProcessor() (HttpRequestProcessor, error)
	HttpRequestProcessorDone(requestProcessor HttpRequestProcessor)
}
