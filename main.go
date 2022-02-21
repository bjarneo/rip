package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/bjarneo/rip/gui"
	"github.com/bjarneo/rip/statistics"
	"github.com/bjarneo/rip/utils"
	"github.com/pterm/pterm"
)

// Initialize our statistics
var stats statistics.Statistics = statistics.NewStatistics()

/*
	If you for some reason end up in a situation where you get
	this error message: "socket: too many open files"

	try to set ulimit to a higher number.
	$ ulimit -n 12000
*/
func request(url string) bool {
	start := utils.NowUnixMilli()

	stats.SetTotal(1)

	resp, err := http.Get(url)

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		stats.SetFailure(1)

		return false
	}

	stats.SetDataTransferred(len(body))

	stats.SetSuccessful(1)

	stop := utils.NowUnixMilli()

	// Update all of our time statistics
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

	spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Load testing %s", url))

	var wg sync.WaitGroup

	// Start the wait groups
	for i := 0; i < concurrent; i++ {
		wg.Add(1)

		// run the concurrent go routines
		go func() {
			for {
				request(url)
			}
		}()
	}

	// This loop will run until the end is reached
	// then it will close the wait groups and break the loop
	for {
		// Run the for loop once a second
		time.Sleep(time.Second * time.Duration(1))

		if utils.NowUnixMilli() < end {
			continue
		}

		// Close all the concurrent wait groups
		for i := 0; i < concurrent; i++ {
			wg.Done()
		}

		break
	}

	// Block until wait groups has been closed
	wg.Wait()

	spinner.Success()

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

	// Run until the interval is done
	workers(*concurrent, *interval, url)

	gui.PrintTable(stats, url)
}
