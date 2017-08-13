package httpServer

import "regexp"
import "errors"

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
			previousKey = split[0]
			headerMap[previousKey] = split[1]
		}
	}

	return headerMap, nil
}
