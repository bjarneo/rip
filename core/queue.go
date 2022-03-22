package core

import (
	"sync"
)

var queueMutex sync.Mutex

type queue struct {
	requests []int
	limit    int
}

func NewQueue(limit int) queue {
	q := queue{
		requests: make([]int, 0),
		limit:    limit,
	}

	return q
}

func (q *queue) Push() int {
	queueMutex.Lock()
	defer queueMutex.Unlock()

	length := len(q.requests)

	// Set a hard limit of 10 (0-9) entries
	if length == q.limit {
		return length
	}

	q.requests = append(q.requests, 1)

	return len(q.requests)
}

func (q *queue) Pop() {
	queueMutex.Lock()
	defer queueMutex.Unlock()

	q.requests = q.requests[:len(q.requests)-1]
}

func (q *queue) Length() int {
	return len(q.requests)
}
