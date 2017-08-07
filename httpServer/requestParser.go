package httpServer

import "strings"
import "strconv"
import "errors"
import "net/url"

func parseRequest(requestStr string) (Request, error) {
	lines := strings.Split(requestStr, CRLF)
	linesLength := len(lines)
	if linesLength == 0 {
		return Request{}, errors.New("No data in request input")
	}

	request, err := parseRequestLine(lines[0])
	if err != nil {
		return request, err
	}

	if request.Version > 1.0 {
		sliceIndex := getLastHeaderIndex(lines)
		request.Headers = parseHeaders(lines[1:sliceIndex])

		if sliceIndex < linesLength {
			request.Body = parseBody(lines[sliceIndex+1 : linesLength])
		}
	}

	return request, err
}

func parseRequestLine(requestLine string) (Request, error) {
	words := strings.Split(requestLine, " ")
	//For http/0.9 it will only be of length 2

	switch length := len(words); length {
	case 3:
		return parseRequestLineV10(words)
	case 2:
		return parseRequestLineV09(words)
	default:
		return Request{}, errors.New("Request line is malformed.")
	}
}

func parseRequestLineV10(words []string) (Request, error) {
	request, err := baseParseRequestLine(words)
	if err != nil {
		return request, err
	}
	request.RequestType = FULL

	return request, nil
}

func parseRequestLineV09(words []string) (Request, error) {
	request, err := baseParseRequestLine(words)
	if err == nil {
		//HTTP/0.9 only supports GET methods
		if request.Method != GET {
			return request, errors.New("HTTP/0.9 only supports GET method.")
		}
		request.RequestType = SIMPLE
	}

	return request, err
}

func baseParseRequestLine(words []string) (Request, error) {
	var request Request

	method, err := getMethod(words[0])
	if err != nil {
		return request, err
	}
	request.Method = method
	request.MethodString = words[0]

	uri, err := getURI(words[1])
	if err != nil {
		return request, err
	}

	request.Path = uri

	if len(words) == 2 {
		request.Version = 0.9
		return request, nil
	}

	if version, err := getHTTPVersion(words[2]); err != nil {
		return request, err
	} else {
		request.Version = version
	}

	return request, nil
}

func getMethod(value string) (Method_t, error) {
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

	versionStr := strings.Trim(value, HTTP)
	version, err := strconv.ParseFloat(versionStr, 32)
	if err != nil {
		return -1, err
	}

	return float32(version), nil
}
