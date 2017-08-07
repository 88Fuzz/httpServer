package httpServer

func processRequest(requestProcessorProvider HttpRequestProcessorProvider, request Request) Response {
	processor, err := requestProcessorProvider.GetHttpRequestProcessor()
	if processor == nil || err != nil {
		var response Response
		response.StatusCode = INTERNAL_SERVER_ERROR

		return response
	}
	defer requestProcessorProvider.HttpRequestProcessorDone(processor)

	switch request.Method {
	case GET:
		return processor.Get(request)
	case HEAD:
		return processor.Head(request)
	case POST:
		return processor.Post(request)
	case EXTENSION:
		return processor.ExtensionMethod(request)
	case INVALID:
		fallthrough
	default:
		var response Response
		response.StatusCode = NOT_IMPLEMENTED
		return response
	}
}
