package core

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

var (
	payloadData []byte = []byte("")
	doOnce      sync.Once
)

type Arguments struct {
	concurrent *int
	interval   *int
	logger     *bool
	host       *string
	hosts      *string
	udp        *bool
	bytes      *int
	post       *bool
	put        *bool
	patch      *bool
	json       *string
	headers    *string
	proxy      *string
	requests   *int
}

func NewArgs() Arguments {
	flags := Arguments{
		concurrent: flag.Int("concurrent", 10, "How many concurrent users to simulate"),
		interval:   flag.Int("interval", 60, "How many seconds to run the test"),
		logger:     flag.Bool("logger", false, "Log the requests to $HOME/rip.log"),
		hosts:      flag.String("hosts", "", "A file of hosts. Each host should be on a new line. It will randomly choose a host."),
		udp:        flag.Bool("udp", false, "Run requests UDP flood attack and not http requests"),
		bytes:      flag.Int("udp-bytes", 2048, "Set the x bytes for the UDP flood attack"),
		post:       flag.Bool("post", false, "POST HTTP request"),
		put:        flag.Bool("put", false, "PUT HTTP request"),
		patch:      flag.Bool("patch", false, "PATCH HTTP request"),
		json:       flag.String("json", "", "Path to the JSON payload file to be used for the HTTP requests"),
		headers:    flag.String("headers", "", "Path to the headers file"),
		proxy:      flag.String("proxy", "", "The proxy URL to route the traffic"),
		requests:   flag.Int("requests", 0, "Max requests per concurrent user at a time"),
	}

	flag.Parse()

	// The host you want to load test
	host := flag.Arg(0)
	if host == "" && *flags.hosts == "" {
		fmt.Print("No host provided. Example: $ rip https://www.google.com, or $ rip -u hosts.txt.")

		os.Exit(1)
	}

	flags.host = &host

	return flags
}

func (flags *Arguments) Concurrent() int {
	return *flags.concurrent
}

func (flags *Arguments) Interval() int {
	return *flags.interval
}

func (flags *Arguments) Logger() bool {
	return *flags.logger
}

func (flags *Arguments) Hosts() []string {
	if *flags.hosts != "" {
		return LinesFromFile(*flags.hosts)
	}

	host := make([]string, 1)

	host[0] = *flags.host

	return host
}

func (flags *Arguments) RequestType() string {
	if *flags.udp {
		return "udp"
	}

	return "http"
}

func (flags *Arguments) Bytes() int {
	return *flags.bytes
}

func (flags *Arguments) HTTPMethod() string {
	if *flags.post {
		return "POST"
	} else if *flags.put {
		return "PUT"
	} else if *flags.patch {
		return "PATCH"
	} else {
		return "GET"
	}
}

func (flags *Arguments) IsJSONPayload() bool {
	return *flags.json != ""
}

func (flags *Arguments) JSONPayload() []byte {
	doOnce.Do(func() {
		if *flags.json != "" {
			payloadData = []byte(FileContent(*flags.json))
		}
	})

	return payloadData
}

func (flags *Arguments) Headers() []string {
	if *flags.headers != "" {
		return LinesFromFile(*flags.headers)
	}

	return make([]string, 0)
}

func (flags *Arguments) Proxy() string {
	return *flags.proxy
}

func (flags *Arguments) Requests() int {
	return *flags.requests
}
