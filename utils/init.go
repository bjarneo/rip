package utils

import (
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

func FileURL(urls string) []string {
	data, err := os.ReadFile(urls)

	if err != nil {
		panic(err)
	}

	urlsToslice := deleteEmptyFromSlice(strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n"))

	return urlsToslice
}
