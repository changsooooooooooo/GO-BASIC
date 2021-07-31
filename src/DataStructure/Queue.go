package main

import (
	"container/list"
	"fmt"
)

const M = 10

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

func formMap() map[string]string {
	m := make(map[string]string, 5)

	m["1"] = "seoul"
	m["2"] = "seoul"
	m["3"] = "seoul"
	m["4"] = "seoul"
	m["5"] = "seoul"
	m["6"] = "seoul"
	m["7"] = "seoul"
	m["8"] = "seoul"
	m["9"] = "seoul"
	m["0"] = "seoul"

	return m
}

func formArr(arr [M]int) {
	arr[2] = 10
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
