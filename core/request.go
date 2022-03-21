package core

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/dchest/uniuri"
)

// Initialize the logger
var l logger = NewLogger()

func udpRequests(hosts []string, args Arguments, stats Statistics) bool {
	bytes := args.Bytes()

	host := RandomSlice(hosts)

	// Never reuse the connection as we want to do a requests towards a random host
	conn, err := net.Dial("udp", host)

	if err != nil {
		return false
	}

	floodString := uniuri.NewLen(bytes)

	conn.Write([]byte(floodString))

	stats.SetDataTransferred(len(floodString))

	if args.Logger() {
		l.Add(host)
	}

	// close the connection as we do not reuse it
	conn.Close()

	return true
}

func httpRequests(hosts []string, args Arguments, stats Statistics) bool {
	host := RandomSlice(hosts)

	req, err := http.NewRequest(args.HTTPMethod(), host, bytes.NewBuffer(args.JSONPayload()))

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	headers := ParseHeaders(args.Headers())

	if args.IsJSONPayload() {
		headers.Add("Content-Type", "application/json; charset=UTF-8")
	}

	headers.Add("User-Agent", "Rest In Peace")

	// Iterate through our custom headers, and add them to the request
	for key, value := range headers.Headers() {
		req.Header.Add(key, value)
	}

	client := &http.Client{}

	if args.Proxy() != "" {
		proxyUrl, err := url.Parse(args.Proxy())

		if err != nil {
			panic(err)
		}

		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}

	resp, err := client.Do(req)

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	if args.Logger() {
		l.Add(host)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	stats.SetDataTransferred(len(body) + len(args.JSONPayload()))

	defer resp.Body.Close()

	return true
}

func Request(hosts []string, args Arguments, stats Statistics) {
	queue := NewQueue(args.Requests())

	go func() {
		defer l.Close()

		for {
			// If we limit the requests to x per concurrent user,
			// run the queue logic
			if args.Requests() > 0 {
				if queue.Length() == args.Requests() {
					continue
				}
				queue.Push()
			}

			start := NowUnixMilli()

			stats.SetTotal(1)

			if args.RequestType() == "http" {
				httpRequests(hosts, args, stats)
			} else if args.RequestType() == "udp" {
				udpRequests(hosts, args, stats)
			}

			stats.SetSuccessful(1)

			stop := NowUnixMilli()

			// Update all of our time statistics
			stats.SetResponseTime(stop - start)
			stats.SetShortest(stop - start)
			stats.SetLongest(stop - start)

			// Pop one request from the queue after the request is done
			if args.Requests() > 0 && queue.Length() == args.Requests() {
				queue.Pop()
			}
		}
	}()
}
