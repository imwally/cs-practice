package main

import "fmt"

type Queue struct {
	Storage []interface{}
}

func (q *Queue) Length() int {
	return len(q.Storage)
}

func (q *Queue) Enqueue(v interface{}) {
	q.Storage = append(q.Storage, v)
}

func (q *Queue) Dequeue() interface{} {
	if q.Length() < 1 {
		return nil
	}

	first := q.Storage[0]
	q.Storage = q.Storage[1:]

	return first
}

func (q Queue) String() string {
	if q.Length() < 1 {
		return ""
	}

	var str string
	for i := len(q.Storage) - 1; i >= 0; i-- {
		str += fmt.Sprintf("%v, ", q.Storage[i])
	}

	return str[:len(str)-2]
}

func main() {
	q := new(Queue)
	q.Enqueue(1)
	fmt.Println(q)
	q.Enqueue(2)
	fmt.Println(q)
	q.Enqueue(3)
	fmt.Println(q)
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
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
