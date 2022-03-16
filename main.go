package main

import (
	"sync"
	"time"

	"github.com/bjarneo/rip/core"
	"github.com/pterm/pterm"
)

// Initialize the cli arguments
var args core.Arguments = core.NewArgs()

// Initialize our statistics
var stats core.TotalStatistics = core.NewStatistics()

func workers(concurrent int, interval int, hosts []string) {
	// Let us start the timer for how long the workers are running
	start := core.NowUnixMilli()
	end := core.FutureUnixMilli(interval)

	area, _ := pterm.DefaultArea.Start()

	var wg sync.WaitGroup

	// Start the wait groups
	for i := 0; i < concurrent; i++ {
		wg.Add(1)

		core.Request(hosts, args, &stats)
	}

	// This loop will run until the end is reached
	// then it will close the wait groups and break the loop
	for {
		// Run the for loop once a second
		time.Sleep(time.Second * time.Duration(1))

		area.Update(core.Logo() + core.PrintStats(&stats))

		if core.NowUnixMilli() < end {
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
	stop := core.NowUnixMilli()

	stats.SetElapsedTime(stop - start)

	// Final update
	area.Update(core.Logo() + core.PrintStats(&stats))

	// stop the area update
	area.Stop()
}

func main() {
	// Run until the interval is done
	workers(args.Concurrent(), args.Interval(), args.Hosts())
}
