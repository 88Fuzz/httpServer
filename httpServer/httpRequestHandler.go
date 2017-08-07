package httpServer

func processRequest(requestProcessorProvider HttpRequestProcessorProvider, req Request) Response {
	processor, err := requestProcessorProvider.GetHttpRequestProcessor()
	if processor == nil || err != nil {
		var response Response
		response.statusCode = INTERNAL_SERVER_ERROR

		return response
	}

	switch req.method {
	case GET:
		return processor.Get(req)
	case HEAD:
		return processor.Head(req)
	case POST:
		return processor.Post(req)
	case EXTENSION:
		return processor.ExtensionMethod(req)
	case INVALID:
		fallthrough
	default:
		var response Response
		response.statusCode = NOT_IMPLEMENTED
		return response
	}
}
