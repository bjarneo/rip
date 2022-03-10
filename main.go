package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bjarneo/rip/request"
	"github.com/bjarneo/rip/statistics"
	"github.com/bjarneo/rip/tui"
	"github.com/bjarneo/rip/utils"
	"github.com/pterm/pterm"
)

// Initialize the cli arguments
var args utils.Arguments = utils.Args()

// Initialize our statistics
var stats statistics.TotalStatistics = statistics.NewStatistics()

func workers(concurrent int, interval int, hosts []string) {
	// Let us start the timer for how long the workers are running
	start := utils.NowUnixMilli()
	end := utils.FutureUnixMilli(interval)

	area, _ := pterm.DefaultArea.Start()

	var wg sync.WaitGroup

	// Start the wait groups
	for i := 0; i < concurrent; i++ {
		wg.Add(1)

		request.Request(hosts, args, &stats)
	}

	// This loop will run until the end is reached
	// then it will close the wait groups and break the loop
	for {
		// Run the for loop once a second
		time.Sleep(time.Second * time.Duration(1))

		area.Update(tui.Logo() + tui.PrintStats(&stats))

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

	// End the timer for how long the workers are running
	stop := utils.NowUnixMilli()

	stats.SetElapsedTime(stop - start)

	// Final update
	area.Update(tui.Logo() + tui.PrintStats(&stats))

	// stop the area update
	area.Stop()
}

func main() {
	// Run until the interval is done
	workers(args.Concurrent(), args.Interval(), args.Hosts())

	if args.Output() == "json" {
		fmt.Println(stats.ToJSON())
	}
}
