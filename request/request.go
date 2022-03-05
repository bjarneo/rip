package request

import (
	"io"
	"math/rand"
	"net"
	"net/http"

	"github.com/bjarneo/rip/statistics"
	"github.com/bjarneo/rip/utils"
	"github.com/dchest/uniuri"
)

// Initialize the logger
var logToFile = utils.Logger()

func udpRequest(url string, args utils.Arguments, stats statistics.Statistics) bool {
	const BYTES = 2048

	conn, err := net.Dial("udp", url)

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	conn.Write([]byte(uniuri.NewLen(BYTES)))

	//conn.Close()

	return true
}

func httpRequest(url string, args utils.Arguments, stats statistics.Statistics) bool {
	resp, err := http.Get(url)

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	if args.Logger() {
		logToFile(url)
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

func Request(urls []string, args utils.Arguments, stats statistics.Statistics) bool {
	start := utils.NowUnixMilli()

	stats.SetTotal(1)

	url := urls[rand.Intn(len(urls))]

	var success bool

	if args.RequestType() == "http" {
		success = httpRequest(url, args, stats)
	} else if args.RequestType() == "udp" {
		success = udpRequest(url, args, stats)
	}

	if !success {
		return false
	}

	stats.SetSuccessful(1)

	stop := utils.NowUnixMilli()

	// Update all of our time statistics
	stats.SetResponseTime(stop - start)
	stats.SetShortest(stop - start)
	stats.SetLongest(stop - start)

	return true
}
