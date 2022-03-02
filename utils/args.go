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
	urls       *string
}

func Args() Arguments {
	flags := Arguments{
		concurrent: flag.Int("c", 10, "How many concurrent users to simulate"),
		interval:   flag.Int("t", 60, "How many seconds to run the test"),
		logger:     flag.Bool("l", false, "Log the requests to $HOME/rip.log"),
		urls:       flag.String("u", "", "A file of URLs. Each URL should be on a new line. It will randomly choose a URL."),
	}

	flag.Parse()

	// The URL you want to load test
	url := flag.Arg(0)
	if url == "" && *flags.urls == "" {
		fmt.Print("No URL provided. Example: $ rip https://www.google.com, or $ rip -u urls.txt.")

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

func (flags *Arguments) Urls() []string {
	if *flags.urls != "" {
		return FileURL(*flags.urls)
	}

	url := make([]string, 1)

	url[0] = *flags.url

	return url
}
