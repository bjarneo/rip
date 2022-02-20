package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	"github.com/bjarneo/rip/statistics"
	"github.com/bjarneo/rip/utils"
	"github.com/pterm/pterm"
)

var stats statistics.Statistics = statistics.NewStatistics()

// ulimit -n 12000
// socket: too many open files
func fetch(url string) {
	stats.SetTotal(1)

	resp, err := http.Get(url)

	if err != nil {
		stats.SetFailed(1)
	}

	stats.SetSuccessful(1)

	defer resp.Body.Close()

	/*
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}
	*/
}

func workers(concurrent int, interval int, url string) {
	end := utils.FutureUnixMilli(interval)

	var wg sync.WaitGroup

	for j := 0; j < concurrent; j++ {
		wg.Add(1)

		go func() {
			for {
				if utils.NowUnixMilli() >= end {
					wg.Done()
					break
				}

				fetch(url)
			}
		}()
	}

	wg.Wait()

	fmt.Println(stats)
}

func main() {
	url := flag.String("url", "", "The url you want to load test")
	concurrent := flag.Int("concurrent", 10, "How many concurrent users to simulate")
	interval := flag.Int("interval", 60, "How many seconds to run the test")

	flag.Parse()

	spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Load testing %s", *url))

	// Run until the interval is done
	workers(*concurrent, *interval, *url)

	spinner.Success()

	pterm.Success.Println("Done")
}
