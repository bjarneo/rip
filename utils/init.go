package utils

import (
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
