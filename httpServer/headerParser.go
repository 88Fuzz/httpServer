package httpServer

import "regexp"
import "errors"
import "strings"

const defaultPreviousKey = "NONE"

var headerDelimiterPattern = regexp.MustCompile(":\\s")

func getLastHeaderIndex(lines []string) int {
	length := len(lines)
	index := 1
	for i := index; i < length; i++ {
		line := lines[i]
		index = i
		if len(line) == 0 {
			//The HTTP spec says a blank line seperates headers from the body
			break
		}
	}

	return index
}

func parseHeaders(lines []string) (map[string]string, error) {
	headerMap := make(map[string]string, len(lines))
	previousKey := defaultPreviousKey
	for _, line := range lines {
		split := headerDelimiterPattern.Split(line, -1)
		if len(split) == 1 {
			//If there's no header key on this line, it means the values should be appending to the previous key
			if previousKey == defaultPreviousKey {
				return headerMap, errors.New("Headers malformed")
			}
			previousValue := headerMap[previousKey]
			headerMap[previousKey] = previousValue + HEADER_VALUE_DELIMITER + split[0]
		} else if len(split) != 2 {
			//Malformed header
			return headerMap, errors.New("Headers malformed")
			continue
		} else {
			previousKey = getStandardHeader(split[0])
			headerMap[previousKey] = split[1]
		}
	}

	return headerMap, nil
}

func getStandardHeader(header string) string {
	switch {
	case strings.EqualFold(header, ACCEPT):
		return ACCEPT
	case strings.EqualFold(header, ACCEPT_CHARSET):
		return ACCEPT_CHARSET
	case strings.EqualFold(header, ACCEPT_ENCODING):
		return ACCEPT_ENCODING
	case strings.EqualFold(header, ACCEPT_LANGUAGE):
		return ACCEPT_LANGUAGE
	case strings.EqualFold(header, ACCEPT_DATETIME):
		return ACCEPT_DATETIME
	case strings.EqualFold(header, ACCESS_CONTROL_REQUEST_METHOD):
		return ACCESS_CONTROL_REQUEST_METHOD
	case strings.EqualFold(header, ACCESS_CONTROL_REQUEST_HEADERS):
		return ACCESS_CONTROL_REQUEST_HEADERS
	case strings.EqualFold(header, AUTHORIZATION):
		return AUTHORIZATION
	case strings.EqualFold(header, CACHE_CONTORL):
		return CACHE_CONTORL
	case strings.EqualFold(header, CONNECTION):
		return CONNECTION
	case strings.EqualFold(header, COOKIE):
		return COOKIE
	case strings.EqualFold(header, CONTENT_LENGTH):
		return CONTENT_LENGTH
	case strings.EqualFold(header, CONTENT_MD5):
		return CONTENT_MD5
	case strings.EqualFold(header, CONTENT_TYPE):
		return CONTENT_TYPE
	case strings.EqualFold(header, DATE):
		return DATE
	case strings.EqualFold(header, EXPECT):
		return EXPECT
	case strings.EqualFold(header, FORWARDED):
		return FORWARDED
	case strings.EqualFold(header, FROM):
		return FROM
	case strings.EqualFold(header, HOST):
		return HOST
	case strings.EqualFold(header, IF_MATCH):
		return IF_MATCH
	case strings.EqualFold(header, IF_MODIFIED_SINCE):
		return IF_MODIFIED_SINCE
	case strings.EqualFold(header, IF_NONE_MATCH):
		return IF_NONE_MATCH
	case strings.EqualFold(header, IF_RANGE):
		return IF_RANGE
	case strings.EqualFold(header, IF_UNMODIFIED_SINCE):
		return IF_UNMODIFIED_SINCE
	case strings.EqualFold(header, MAX_FORWARDS):
		return MAX_FORWARDS
	case strings.EqualFold(header, ORIGIN):
		return ORIGIN
	case strings.EqualFold(header, PRAGMA):
		return PRAGMA
	case strings.EqualFold(header, PROXY_AUTHORIZATION):
		return PROXY_AUTHORIZATION
	case strings.EqualFold(header, RANGE):
		return RANGE
	case strings.EqualFold(header, REFERER):
		return REFERER
	case strings.EqualFold(header, TE):
		return TE
	case strings.EqualFold(header, USER_AGENT):
		return USER_AGENT
	case strings.EqualFold(header, UPGRADE):
		return UPGRADE
	case strings.EqualFold(header, VIA):
		return VIA
	case strings.EqualFold(header, WARNING):
		return WARNING
	default:
		return header
	}
}
