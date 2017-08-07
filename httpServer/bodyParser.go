package httpServer

import "bytes"

func parseBody(lines []string) string {
	var bodyBuffer bytes.Buffer
	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		bodyBuffer.WriteString(line + "\r\n")
	}

	return bodyBuffer.String()
}
