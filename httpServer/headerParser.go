package httpServer

import "strings"

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

func parseHeaders(lines []string) map[string]string {
	headerMap := make(map[string]string, len(lines))
	for _, line := range lines {
		split := strings.Split(line, ": ")
		if len(split) != 2 {
			//Malformed header ignore it?
			continue
		}

		headerMap[split[0]] = split[1]
	}

	return headerMap
}
