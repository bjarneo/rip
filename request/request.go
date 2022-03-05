package request

import (
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

	conn.Write([]byte(uniuri.NewLen(bytes)))

	if args.Logger() {
		logToFile(host)
	}

	return true
}

func httpRequests(hosts []string, args utils.Arguments, stats statistics.Statistics) bool {
	host := utils.RandomSlice(hosts)

	resp, err := http.Get(host)

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

	stats.SetDataTransferred(len(body))

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
