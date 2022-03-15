package core

import (
	"fmt"
	"log"
	"os"
	"time"
)

func homeFolder() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return dirname
}

// Logger writes each request to $HOME/rip.log.
func NewLogger() func(string) {
	f, err := os.Create(fmt.Sprintf("%s/rip.log", homeFolder()))
	if err != nil {
		panic(err)
	}

	return func(request string) {
		// use ISO8601 formatting
		logString := fmt.Sprintf(`[%s] - "%s"`, time.Now().UTC().Format("2006-01-02T15:04:05-0700"), request)

		_, err = fmt.Fprintln(f, logString)
		if err != nil {
			panic(err)
		}
	}
}
