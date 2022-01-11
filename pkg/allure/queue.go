package allure

import "sync"

type FiloQueue interface {
	Push(*Step)
	Pop() *Step
	Last() *Step
}

type NestingQueue struct {
	queue []*Step
	count int
	mtx   sync.Mutex
}

func NewNestingQueue() NestingQueue {
	return NestingQueue{queue: make([]*Step, 0)}
}

func (q *NestingQueue) Push(step *Step) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.queue = append(q.queue, step)
	q.count++
}

func (q *NestingQueue) Pop() *Step {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if q.count == 0 {
		return nil
	}
	q.count--
	result := q.queue[q.count]
	q.queue = q.queue[:q.count]
	return result
}

func (q *NestingQueue) Last() *Step {
	if q.count == 0 {
		return nil
	}
	return q.queue[q.count-1]
}
