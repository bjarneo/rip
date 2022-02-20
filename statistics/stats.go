package statistics

type Statistics struct {
	total           int64
	successful      int64
	failed          int64
	longest         float64
	shortest        float64
	elapsedTime     float64
	responseTime    float64
	dataTransferred float64
}

func NewStatistics() Statistics {
	stats := Statistics{
		total:           0,
		successful:      0,
		failed:          0,
		longest:         0.0,
		shortest:        0.0,
		elapsedTime:     0.0,
		responseTime:    0.0,
		dataTransferred: 0.0,
	}

	return stats
}

func (stats *Statistics) SetTotal(total int64) {
	stats.total += total
}

func (stats *Statistics) Total() int64 {
	return stats.total
}

func (stats *Statistics) SetSuccessful(successful int64) {
	stats.successful += successful
}

func (stats *Statistics) Successful() int64 {
	return stats.successful
}

func (stats *Statistics) SetFailed(failed int64) {
	stats.failed += failed
}

func (stats *Statistics) Failed() int64 {
	return stats.failed
}

func (stats *Statistics) SetLongest(longest float64) {
	stats.longest = longest
}

func (stats *Statistics) Longest() float64 {
	return stats.longest
}

func (stats *Statistics) SetShortest(shortest float64) {
	stats.shortest = shortest
}

func (stats *Statistics) Shortest() float64 {
	return stats.shortest
}

func (stats *Statistics) SetElapsedTime(elapsedTime float64) {
	stats.elapsedTime = elapsedTime
}

func (stats *Statistics) ElapsedTime() float64 {
	return stats.elapsedTime
}

func (stats *Statistics) SetResponseTime(responseTime float64) {
	stats.responseTime = responseTime
}

func (stats *Statistics) ResponseTime() float64 {
	return stats.responseTime
}

func (stats *Statistics) SetDataTransferred(dataTransferred float64) {
	stats.dataTransferred += dataTransferred
}

func (stats *Statistics) DataTransferred() float64 {
	return stats.dataTransferred
}
