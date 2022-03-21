package core

import "sync"

var mu sync.Mutex

type Statistics interface {
	SetTotal(totalt int64)
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
	total           int64
	successful      int64
	failed          int64
	longest         int64
	shortest        int64
	elapsedTime     int64
	responseTime    []int64
	dataTransferred int
}

func NewStatistics() TotalStatistics {
	stats := TotalStatistics{
		total:           0,
		successful:      0,
		failed:          0,
		longest:         0,
		shortest:        0,
		elapsedTime:     0,
		responseTime:    make([]int64, 1),
		dataTransferred: 0,
	}

	return stats
}

func (stats *TotalStatistics) SetTotal(total int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.total += total
}

func (stats *TotalStatistics) Total() int64 {
	return stats.total
}

func (stats *TotalStatistics) SetSuccessful(successful int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.successful += successful
}

func (stats *TotalStatistics) Successful() int64 {
	return stats.successful
}

func (stats *TotalStatistics) SetFailure(failed int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.failed += failed
}

func (stats *TotalStatistics) Failure() int64 {
	return stats.failed
}

func (stats *TotalStatistics) SetLongest(longest int64) {
	if longest > stats.longest {
		stats.longest = longest
	}
}

func (stats *TotalStatistics) Longest() int64 {
	return stats.longest
}

func (stats *TotalStatistics) SetShortest(shortest int64) {
	if stats.shortest == 0 {
		stats.shortest = shortest
	}

	if shortest < stats.shortest {
		stats.shortest = shortest
	}
}

func (stats *TotalStatistics) Shortest() int64 {
	return stats.shortest
}

func (stats *TotalStatistics) SetElapsedTime(elapsedTime int64) {
	stats.elapsedTime = elapsedTime
}

func (stats *TotalStatistics) ElapsedTime() int64 {
	return stats.elapsedTime
}

func (stats *TotalStatistics) SetResponseTime(responseTime int64) {
	stats.responseTime = append(stats.responseTime, responseTime)
}

func (stats *TotalStatistics) ResponseTime() int {
	var sum int64 = 0

	for _, res := range stats.responseTime {
		sum += res
	}

	// Calculate the average response time before returning the int
	avg := int(sum) / len(stats.responseTime)

	return avg
}

func (stats *TotalStatistics) SetDataTransferred(dataTransferred int) {
	mu.Lock()
	defer mu.Unlock()

	stats.dataTransferred += dataTransferred
}

func (stats *TotalStatistics) DataTransferred() int {
	return stats.dataTransferred
}
