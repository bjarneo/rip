package gui

import (
	"fmt"

	"github.com/bjarneo/rip/statistics"
	"github.com/pterm/pterm"
)

func PrintTable(stats statistics.Statistics, url string) {
	fmt.Println()

	pterm.DefaultTable.WithHasHeader().WithData(
		pterm.TableData{
			{"URL", "Total", "Successful", "Failed", "Longest", "Shortest", "Elapsed Time", "Avg Response Time", "Data transferred"},
			{
				url,
				fmt.Sprintf("%d", stats.Total()),
				fmt.Sprintf("%d", stats.Successful()),
				fmt.Sprintf("%d", stats.Failure()),
				fmt.Sprintf("%dms", stats.Longest()),
				fmt.Sprintf("%dms", stats.Shortest()),
				fmt.Sprintf("%.2fs", float64(stats.ElapsedTime()/1000)),
				fmt.Sprintf("%.2fms", float64(stats.ResponseTime())),
				fmt.Sprintf("%.2fkb", float64(stats.DataTransferred()/1000)),
			},
		},
	).Render()
}
