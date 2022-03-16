package core

import (
	"fmt"
	"log"
	"os"
	"time"
)

type logger struct {
	file *os.File
}

var logPath string = fmt.Sprintf("%s/rip.log", homeFolder())

func homeFolder() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return dirname
}

// Logger writes each request to $HOME/rip.log.
func NewLogger() logger {
	f, err := os.Create(logPath)
	if err != nil {
		panic(err)
	}

	l := logger{
		file: f,
	}

	return l
}

func (l *logger) Add(request string) {
	// use ISO8601 formatting
	logString := fmt.Sprintf(
		`[%s] - "%s"`,
		time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		request,
	)

	fmt.Fprintln(l.file, logString)
}

func (l *logger) Close() {
	l.file.Close()
}
