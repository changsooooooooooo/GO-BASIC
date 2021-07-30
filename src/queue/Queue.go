package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	queue *list.List
}

func (q *Queue) push(value interface{}) {
	q.queue.PushFront(value)
}

func (q *Queue) pop() interface{} {
	front := q.queue.Front()
	if front == nil {
		return nil
	}
	return q.queue.Remove(front)
}

func newQueue() *Queue {
	return &Queue{list.New()}
	//func New() *List { return new(List).Init() }
	//return Type = *List
}

func main() {
	q := newQueue()

	for i := 0; i < 5; i++ {
		q.push(i)
	}

	v := q.pop()
	for v != nil {
		fmt.Println(v)
		v = q.pop()
	}
}
