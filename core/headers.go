package core

import (
	"regexp"
)

// This will match <first-group>:<second-group> of a header
// Test it here https://regex101.com/r/cn0QvY/1
const (
	PATTERN = "([a-zA-Z0-9-]+)[^*](.*)"
)

// Pre define the compiler
var pattern *regexp.Regexp = regexp.MustCompile(PATTERN)

type headers struct {
	entries map[string]string
}

func ParseHeaders(headersFileContent []string) headers {
	h := headers{
		entries: make(map[string]string),
	}

	for _, line := range headersFileContent {
		header := pattern.FindStringSubmatch(line)

		h.Add(header[1], header[2])
	}

	return h
}

func (h *headers) Add(key string, value string) {
	h.entries[key] = value
}

func (h *headers) Headers() map[string]string {
	return h.entries
}
