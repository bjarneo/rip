package utils

import "strings"

func ParseHeaders(headersFileContent []string) map[string]string {
	headers := make(map[string]string, 0)

	headers["User-Agent"] = "Rest In Peace"

	for _, line := range headersFileContent {
		header := strings.Split(line, ":")

		headers[header[0]] = strings.TrimSpace(header[1])
	}

	return headers
}
