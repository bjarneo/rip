package core

import (
	"sync"
	"time"

	"github.com/pterm/pterm"
)

// Initialize the cli arguments
var args Arguments = NewArgs()

// Initialize our statistics
var stats TotalStatistics = NewStatistics()

func Execute() {
	// Let us start the timer for how long the workers are running
	start := NowUnixMilli()
	end := FutureUnixMilli(args.Interval())

	ParseHeaders(args)

	area, _ := pterm.DefaultArea.Start()

	var wg sync.WaitGroup

	// Start the wait groups
	for i := 0; i < args.Concurrent(); i++ {
		wg.Add(1)

		Request(args.Hosts(), args, &stats)
	}

	// This loop will run until the end is reached
	// then it will close the wait groups and break the loop
	for {
		// Run the for loop once a second
		time.Sleep(time.Second * time.Duration(1))

		stats.SetElapsedTime(NowUnixMilli() - start)

		area.Update(Logo() + PrintStats(&stats))

		if NowUnixMilli() < end {
			continue
		}

		// Close all the concurrent wait groups
		for i := 0; i < args.Concurrent(); i++ {
			wg.Done()
		}

		break
	}

	// Block until wait groups has been closed
	wg.Wait()

	// End the timer for how long the workers are running
	stop := NowUnixMilli()

	stats.SetElapsedTime(stop - start)

	// Final update
	area.Update(Logo() + PrintStats(&stats))

	// stop the area update
	area.Stop()
}
