package core

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func NowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

func FutureUnixMilli(interval int) int64 {
	t := time.Now()

	future := t.Add(time.Second * time.Duration(interval))

	return future.UnixMilli()
}

// copy and pasted from stackoverflow because I am lazy
func deleteEmptyFromSlice(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func FileContent(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	return string(data)
}

func LinesFromFile(file string) []string {
	return deleteEmptyFromSlice(
		strings.Split(
			strings.ReplaceAll(
				FileContent(file),
				"\r\n",
				"\n",
			),
			"\n",
		),
	)
}

func RandomSlice(hosts []string) string {
	return hosts[rand.Intn(len(hosts))]
}
