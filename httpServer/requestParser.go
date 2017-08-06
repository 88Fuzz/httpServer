package httpServer

import "strings"
import "strconv"
import "errors"
import "net/url"

func parseRequest(requestStr string) (request, error) {
	lines := strings.Split(requestStr, "\r\n")
	if len(lines) == 0 {
		return request{}, errors.New("No data in request input")
	}

	req, err := parseRequestLine(lines[0])
	if err != nil {
		return req, err
	}

	if req.version > 1.0 {
		sliceIndex := getLastHeaderIndex(lines)
		headerLines := lines[1:sliceIndex]
		req.headers = parseHeaders(headerLines)
	}

	//TODO parse the rest of the request, just the body is left
	return req, err
}

func parseRequestLine(requestLine string) (request, error) {
	words := strings.Split(requestLine, " ")
	//For http/0.9 it will only be of length 2

	switch length := len(words); length {
	case 3:
		return parseRequestLineV10(words)
	case 2:
		return parseRequestLineV09(words)
	default:
		return request{}, errors.New("Request line is malformed.")
	}
}

func parseRequestLineV10(words []string) (request, error) {
	req, err := baseParseRequestLine(words)
	if err != nil {
		return req, err
	}
	req.requestType = FULL

	return req, nil
}

func parseRequestLineV09(words []string) (request, error) {
	req, err := baseParseRequestLine(words)
	if err == nil {
		//HTTP/0.9 only supports GET methods
		if req.method != GET {
			return req, errors.New("HTTP/0.9 only supports GET method.")
		}
		req.requestType = SIMPLE
	}

	return req, err
}

func baseParseRequestLine(words []string) (request, error) {
	var req request

	method, err := getMethod(words[0])
	if err != nil {
		return req, err
	}
	req.method = method
	req.methodValue = words[0]

	uri, err := getURI(words[1])
	if err != nil {
		return req, err
	}

	req.path = uri

	if len(words) == 2 {
		req.version = 0.9
		return req, nil
	}

	if version, err := getHTTPVersion(words[2]); err != nil {
		return req, err
	} else {
		req.version = version
	}

	return req, nil
}

func getMethod(value string) (METHOD, error) {
	if length := len(value); length == 0 {
		return INVALID, errors.New("Invalid method.")
	}

	switch value {
	case "GET":
		return GET, nil
	case "HEAD":
		return HEAD, nil
	case "POST":
		return POST, nil
	default:
		return EXTENSION, nil
	}
}

func getURI(value string) (string, error) {
	if length := len(value); length == 0 {
		return "", errors.New("Invalid path.")
	}

	return url.QueryUnescape(value)
}

func getHTTPVersion(value string) (float32, error) {
	if length := len(value); length == 0 {
		return -1, errors.New("Invalid HTTP version.")
	}

	versionStr := strings.Trim(value, "HTTP/")
	version, err := strconv.ParseFloat(versionStr, 32)
	if err != nil {
		return -1, err
	}

	return float32(version), nil
}
