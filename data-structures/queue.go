package main

import "fmt"

type Queue struct {
	Storage []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.Storage = append(q.Storage, v)
}

func (q *Queue) Dequeue() interface{} {
	first := q.Storage[0]
	q.Storage = q.Storage[1:]

	return first
}

func main() {
	q := new(Queue)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue("Four")

	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
}
