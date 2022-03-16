package core

import (
	"regexp"
)

// This will match <first-group>:<second-group> of a header
// Test it here https://regex101.com/r/cn0QvY/1
const (
	PATTERN = "([a-zA-Z0-9-]+)[^*](.*)"
)

type headers struct {
	headers map[string]string
}

func ParseHeaders(headersFileContent []string) headers {
	pattern := regexp.MustCompile(PATTERN)

	h := headers{
		headers: make(map[string]string, 0),
	}

	h.add("User-Agent", "Rest In Peace")

	for _, line := range headersFileContent {
		header := pattern.FindStringSubmatch(line)

		h.add(header[1], header[2])
	}

	return h
}

func (h *headers) add(key string, value string) {
	h.headers[key] = value
}

func (h *headers) Headers() map[string]string {
	return h.headers
}
