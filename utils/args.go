package utils

import (
	"flag"
	"fmt"
	"os"
)

type Arguments struct {
	concurrent *int
	interval   *int
	logger     *bool
	url        *string
}

func Args() Arguments {
	flags := Arguments{
		concurrent: flag.Int("c", 10, "How many concurrent users to simulate"),
		interval:   flag.Int("t", 60, "How many seconds to run the test"),
		logger:     flag.Bool("l", false, "Log the requests to $HOME/rip.log"),
	}

	flag.Parse()

	// The URL you want to load test
	url := flag.Arg(0)
	if url == "" {
		fmt.Print("No URL provided. Example: $ rip https://www.google.com")

		os.Exit(1)
	}

	flags.url = &url

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

func (flags *Arguments) Url() string {
	return *flags.url
}
