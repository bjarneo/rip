package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/bjarneo/rip/gui"
	"github.com/bjarneo/rip/statistics"
	"github.com/bjarneo/rip/utils"
	"github.com/pterm/pterm"
)

var stats statistics.Statistics = statistics.NewStatistics()

// ulimit -n 12000
// socket: too many open files
func request(url string) bool {
	start := utils.NowUnixMilli()

	stats.SetTotal(1)

	resp, err := http.Get(url)

	if err != nil {
		stats.SetFailed(1)

		return false
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		stats.SetFailed(1)

		return false
	}

	stats.SetDataTransferred(len(body))

	stats.SetSuccessful(1)

	stop := utils.NowUnixMilli()

	stats.SetResponseTime(stop - start)
	stats.SetShortest(stop - start)
	stats.SetLongest(stop - start)

	defer resp.Body.Close()

	return true
}

func workers(concurrent int, interval int, url string) {
	// Let us start the timer for how long the workers are running
	start := utils.NowUnixMilli()

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

				request(url)
			}
		}()
	}

	wg.Wait()

	// End the timer for how long the workers are running
	stop := utils.NowUnixMilli()

	stats.SetElapsedTime(stop - start)
}

func main() {
	// Custom cli flags
	concurrent := flag.Int("concurrent", 10, "How many concurrent users to simulate")
	interval := flag.Int("interval", 60, "How many seconds to run the test")

	flag.Parse()

	// The URL you want to load test
	url := flag.Arg(0)
	if url == "" {
		fmt.Print("No URL provided. Example: $ rip https://www.google.com")

		os.Exit(1)
	}

	fmt.Println(*concurrent)
	fmt.Println(*interval)

	spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Load testing %s", url))

	// Run until the interval is done
	workers(*concurrent, *interval, url)

	spinner.Success()

	pterm.Success.Println("Done")

	gui.PrintTable(stats, url)
}
