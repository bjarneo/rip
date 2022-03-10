package statistics

import (
	"encoding/json"
	"sync"
)

var mu sync.Mutex

type Statistics interface {
	SetTotal(total int64)
	Total() int64
	SetSuccessful(successful int64)
	Successful() int64
	SetFailure(failed int64)
	Failure() int64
	SetShortest(shortest int64)
	Shortest() int64
	SetLongest(longest int64)
	Longest() int64
	SetElapsedTime(elapsedTime int64)
	ElapsedTime() int64
	SetResponseTime(responseTime int64)
	ResponseTime() int
	SetDataTransferred(dataTransferred int)
	DataTransferred() int
}

type TotalStatistics struct {
	TotalReq             int64   `json:"total"`
	SuccessfulReq        int64   `json:"successful"`
	FailedReq            int64   `json:"failed"`
	LongestReq           int64   `json:"longest"`
	ShortestReq          int64   `json:"shortest"`
	TotalElapsedTime     int64   `json:"elapsedTime"`
	AvgResponseTime      []int64 `json:"responseTime"`
	TotalDataTransferred int     `json:"dataTransferred"`
}

func NewStatistics() TotalStatistics {
	stats := TotalStatistics{
		TotalReq:             0,
		SuccessfulReq:        0,
		FailedReq:            0,
		LongestReq:           0,
		ShortestReq:          0,
		TotalElapsedTime:     0,
		AvgResponseTime:      make([]int64, 1, 1),
		TotalDataTransferred: 0,
	}

	return stats
}

func (stats *TotalStatistics) SetTotal(total int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.TotalReq += total
}

func (stats *TotalStatistics) Total() int64 {
	return stats.TotalReq
}

func (stats *TotalStatistics) SetSuccessful(successful int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.SuccessfulReq += successful
}

func (stats *TotalStatistics) Successful() int64 {
	return stats.SuccessfulReq
}

func (stats *TotalStatistics) SetFailure(failed int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.FailedReq += failed
}

func (stats *TotalStatistics) Failure() int64 {
	return stats.FailedReq
}

func (stats *TotalStatistics) SetLongest(longest int64) {
	if longest > stats.LongestReq {
		stats.LongestReq = longest
	}
}

func (stats *TotalStatistics) Longest() int64 {
	return stats.LongestReq
}

func (stats *TotalStatistics) SetShortest(shortest int64) {
	if stats.ShortestReq == 0 {
		stats.ShortestReq = shortest
	}

	if shortest < stats.ShortestReq {
		stats.ShortestReq = shortest
	}
}

func (stats *TotalStatistics) Shortest() int64 {
	return stats.ShortestReq
}

func (stats *TotalStatistics) SetElapsedTime(elapsedTime int64) {
	stats.TotalElapsedTime = elapsedTime
}

func (stats *TotalStatistics) ElapsedTime() int64 {
	return stats.TotalElapsedTime
}

func (stats *TotalStatistics) SetResponseTime(responseTime int64) {
	stats.AvgResponseTime = append(stats.AvgResponseTime, responseTime)
}

func (stats *TotalStatistics) ResponseTime() int {
	var sum int64 = 0

	for _, res := range stats.AvgResponseTime {
		sum += res
	}

	// Calculate the average response time before returning the int
	avg := int(sum) / len(stats.AvgResponseTime)

	return avg
}

func (stats *TotalStatistics) SetDataTransferred(dataTransferred int) {
	mu.Lock()
	defer mu.Unlock()

	stats.TotalDataTransferred += dataTransferred
}

func (stats *TotalStatistics) DataTransferred() int {
	return stats.TotalDataTransferred
}

func (stats *TotalStatistics) ToJSON() string {
	json, err := json.Marshal(stats)

	if err != nil {
		panic(err)
	}

	return string(json)
}
