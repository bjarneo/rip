package core

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
)

func template(name string, value string) string {
	const MAX_LENGTH int = 42
	pad := MAX_LENGTH - len(name)

	// Example output
	// [ Total                                  1883 ]
	return fmt.Sprintf("[ %s %*s ]", name, pad, value)
}

func Logo() string {
	logo := `

  ___ ___ ___
 | _ \_ _| _ \
 |   /| ||  _/
 |_|_\___|_|

`

	return pterm.DefaultCenter.Sprint(logo)
}

func PrintStats(stats Statistics) string {
	out := []string{
		template("Total", fmt.Sprintf("%d", stats.Total())),
		template("Successful", fmt.Sprintf("%d", stats.Successful())),
		template("Failed", fmt.Sprintf("%d", stats.Failure())),
		template("Unfinished", fmt.Sprintf("%d", stats.Total()-stats.Successful()-stats.Failure())),
		template("Longest", fmt.Sprintf("%dms", stats.Longest())),
		template("Shortest", fmt.Sprintf("%dms", stats.Shortest())),
		template("Elapsed Time", fmt.Sprintf("%.2fs", float64(stats.ElapsedTime()/1000))),
		template("Avg Response Time", fmt.Sprintf("%.2fms", float64(stats.ResponseTime()))),
		template("Data transferred", fmt.Sprintf("%.2fkb", float64(stats.DataTransferred()/1000))),
	}

	return pterm.DefaultCenter.Sprint(strings.Join(out, "\n"))
}
