package protocol

import (
	"sync"
)

type Queue struct {
	items []interface{}
	lock  sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items) == 0
}

func (q *Queue) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items)
}
