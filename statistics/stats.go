package statistics

import "sync"

var mu sync.Mutex

type Statistics struct {
	total           int64
	successful      int64
	failed          int64
	longest         int64
	shortest        int64
	elapsedTime     int64
	responseTime    []int64
	dataTransferred int
}

func NewStatistics() Statistics {
	stats := Statistics{
		total:           0,
		successful:      0,
		failed:          0,
		longest:         0,
		shortest:        0,
		elapsedTime:     0,
		responseTime:    make([]int64, 1, 1),
		dataTransferred: 0,
	}

	return stats
}

func (stats *Statistics) SetTotal(total int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.total += total
}

func (stats *Statistics) Total() int64 {
	return stats.total
}

func (stats *Statistics) SetSuccessful(successful int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.successful += successful
}

func (stats *Statistics) Successful() int64 {
	return stats.successful
}

func (stats *Statistics) SetFailure(failed int64) {
	mu.Lock()
	defer mu.Unlock()

	stats.failed += failed
}

func (stats *Statistics) Failure() int64 {
	return stats.failed
}

func (stats *Statistics) SetLongest(longest int64) {
	if longest > stats.longest {
		stats.longest = longest
	}
}

func (stats *Statistics) Longest() int64 {
	return stats.longest
}

func (stats *Statistics) SetShortest(shortest int64) {
	if stats.shortest == 0 {
		stats.shortest = shortest
	}

	if shortest < stats.shortest {
		stats.shortest = shortest
	}
}

func (stats *Statistics) Shortest() int64 {
	return stats.shortest
}

func (stats *Statistics) SetElapsedTime(elapsedTime int64) {
	stats.elapsedTime = elapsedTime
}

func (stats *Statistics) ElapsedTime() int64 {
	return stats.elapsedTime
}

func (stats *Statistics) SetResponseTime(responseTime int64) {
	stats.responseTime = append(stats.responseTime, responseTime)
}

func (stats *Statistics) ResponseTime() int {
	var sum int64 = 0

	for _, res := range stats.responseTime {
		sum += res
	}

	// Calculate the average response time before returning the int
	avg := int(sum) / len(stats.responseTime)

	return avg
}

func (stats *Statistics) SetDataTransferred(dataTransferred int) {
	mu.Lock()
	defer mu.Unlock()

	stats.dataTransferred += dataTransferred
}

func (stats *Statistics) DataTransferred() int {
	return stats.dataTransferred
}
