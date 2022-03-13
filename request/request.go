package request

import (
	"bytes"
	"io"
	"net"
	"net/http"

	"github.com/bjarneo/rip/statistics"
	"github.com/bjarneo/rip/utils"
	"github.com/dchest/uniuri"
)

// Initialize the logger
var logToFile = utils.Logger()

func udpRequests(hosts []string, args utils.Arguments, stats statistics.Statistics) bool {
	bytes := args.Bytes()

	host := utils.RandomSlice(hosts)

	// Never reuse the connection as we want to do a requests towards a random host
	conn, err := net.Dial("udp", host)

	if err != nil {
		return false
	}

	floodString := uniuri.NewLen(bytes)

	conn.Write([]byte(floodString))

	stats.SetDataTransferred(len(floodString))

	if args.Logger() {
		logToFile(host)
	}

	// close the connection as we do not reuse it
	conn.Close()

	return true
}

func httpRequests(hosts []string, args utils.Arguments, stats statistics.Statistics) bool {
	host := utils.RandomSlice(hosts)

	req, err := http.NewRequest(args.HTTPMethod(), host, bytes.NewBuffer(args.JSONPayload()))

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	headers := utils.ParseHeaders(args.Headers())

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	if args.IsJSONPayload() {
		req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	if args.Logger() {
		logToFile(host)
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

func Request(hosts []string, args utils.Arguments, stats statistics.Statistics) {
	go func() {
		for {
			start := utils.NowUnixMilli()

			stats.SetTotal(1)

			if args.RequestType() == "http" {
				httpRequests(hosts, args, stats)
			} else if args.RequestType() == "udp" {
				udpRequests(hosts, args, stats)
			}

			stats.SetSuccessful(1)

			stop := utils.NowUnixMilli()

			// Update all of our time statistics
			stats.SetResponseTime(stop - start)
			stats.SetShortest(stop - start)
			stats.SetLongest(stop - start)
		}
	}()
}
