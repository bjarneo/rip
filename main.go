package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	"github.com/pterm/pterm"
)

// ulimit -n 12000
// socket: too many open files
func request(url string) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	/*
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}
	*/
}

func worker(wg *sync.WaitGroup, url string) func() chan string {
	ch := make(chan string)

	// Return as closure, and execute the goroutine at a later stage
	return func() chan string {
		// Here goes all the goroutine logic
		go func() {
			defer wg.Done()
			defer close(ch)

			request(url)
		}()

		return ch
	}
}

func main() {
	url := flag.String("url", "", "The url you want to load test")
	concurrent := flag.Int("concurrent", 10, "How many concurrent users to simulate")
	requests := flag.Int("requests", 1, "How many requests per concurrent user")

	flag.Parse()

	p, _ := pterm.DefaultProgressbar.WithTotal(*concurrent * *requests).WithTitle(fmt.Sprintf("Load testing %s", *url)).Start()

	var wg sync.WaitGroup

	var deferred []func() chan string

	/*
		Iterate through all the concurrent users and requests.
		1. Create a channel per request
		2. Add the request to a wait group
		3. Append the worker to the deferred channel slice

		However, what we really want is to create *concurrent
		worker groups, then execute the requests concurrently.
	*/
	for j := 0; j < *requests**concurrent; j++ {
		wg.Add(1)
		deferred = append(deferred, worker(&wg, *url))
	}

	// Now we execute all the channels with the workers
	for _, execWorker := range deferred {
		<-execWorker()

		p.Increment()
	}

	wg.Wait()

	pterm.Success.Println("Done")
}
