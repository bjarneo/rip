package main

import (
	"github.com/bjarneo/rip/core"
)

func main() {
	// Initialize the logger
	logger := core.NewLogger()
	defer logger.Close()

	// Run until the interval is done
	core.Execute()
}
