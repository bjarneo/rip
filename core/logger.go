package core

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type logger struct {
	file *os.File
	mu   sync.Mutex
}

var logPath string = fmt.Sprintf("%s/rip.log", homeFolder())

var loggerInstance *logger

func homeFolder() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return dirname
}

func headerString() string {
	headers := HeaderInstance()
	asString := ""

	// Iterate through our custom headers, and add them to the request
	for key, value := range headers.Headers() {
		asString += fmt.Sprintf(` "%s: %s"`, key, value)
	}

	return asString
}

// Logger writes each request to $HOME/rip.log.
func NewLogger() *logger {
	var once sync.Once

	once.Do(func() {
		f, err := os.Create(logPath)
		if err != nil {
			panic(err)
		}

		loggerInstance = &logger{
			file: f,
		}
	})

	return loggerInstance
}

func Logger() *logger {
	return loggerInstance
}

func (l *logger) Add(request string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// use ISO8601 formatting
	logString := fmt.Sprintf(
		`[%s] - "%s" -%s`,
		time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		request,
		headerString(),
	)

	fmt.Fprintln(l.file, logString)
}

func (l *logger) Close() {
	l.file.Close()
}
