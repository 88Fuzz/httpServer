package httpServer

import "bytes"
import "strconv"

var statusStrings = map[STATUS_CODE]string{
	OK:                    "OK",
	CREATED:               "Created",
	ACCEPTED:              "Accepted",
	NO_CONTENT:            "No Content",
	MOVED_PERMANENTLY:     "Moved Permanently",
	MOVED_TEMPORARILY:     "Moved Temporarily",
	NOT_MODIFIED:          "Not Modified",
	BAD_REQUEST:           "Bad Request",
	UNAUTHORIZED:          "Unauthorized",
	FORBIDDEN:             "Forbidden",
	NOT_FOUND:             "Not Found",
	INTERNAL_SERVER_ERROR: "Server Error",
	NOT_IMPLEMENTED:       "Not Implemented",
	BAD_GATEWAY:           "Bad Gateway",
	SERVICE_UNAVAILABLE:   "Service Unavailable",
}

var statusValues = map[STATUS_CODE]int{
	OK:                    200,
	CREATED:               201,
	ACCEPTED:              202,
	NO_CONTENT:            204,
	MOVED_PERMANENTLY:     301,
	MOVED_TEMPORARILY:     302,
	NOT_MODIFIED:          304,
	BAD_REQUEST:           400,
	UNAUTHORIZED:          401,
	FORBIDDEN:             403,
	NOT_FOUND:             404,
	INTERNAL_SERVER_ERROR: 500,
	NOT_IMPLEMENTED:       501,
	BAD_GATEWAY:           501,
	SERVICE_UNAVAILABLE:   503,
}

func createHttpResponse(request Request, response Response) string {
	var buffer bytes.Buffer
	//Simple requests, i.e. HTTP/0.9 only supports returning the body
	if request.requestType == FULL {
		buffer.WriteString(createStatusLine(response))
		buffer.WriteString(createHeaders(response))
	}

	if request.method != HEAD {
		buffer.WriteString(createBody(response))
	}

	return buffer.String()
}

func createStatusLine(response Response) string {
	statusString := getStatusString(response.statusCode)
	return HTTP + HTTP_VERSION + " " + strconv.Itoa(getStatusValue(response.statusCode)) + " " + statusString + CRLF
}

func createHeaders(response Response) string {
	var buffer bytes.Buffer
	for key, value := range response.headers {
		if key == CONTENT_LENGTH {
			continue
		}
		buffer.WriteString(createHeaderString(key, value))
	}

	if length := len(response.body); length != 0 {
		buffer.WriteString(createHeaderString(CONTENT_LENGTH, strconv.Itoa(length)))
	}

	return buffer.String()
}

func createBody(response Response) string {
	return CRLF + response.body + CRLF
}

func getStatusString(statusCode STATUS_CODE) string {
	return statusStrings[statusCode]
}

func getStatusValue(statusCode STATUS_CODE) int {
	value := statusValues[statusCode]
	if value == 0 {
		return statusValues[INTERNAL_SERVER_ERROR]
	}
	return value
}

func createHeaderString(key string, value string) string {
	return key + HEADER_DELIMITER + value
}
