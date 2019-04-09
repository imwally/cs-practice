package main

import "fmt"

type Stack struct {
	Storage []int
}

func (s *Stack) Push(n int) {
	s.Storage = append(s.Storage, n)
}

func (s *Stack) Pop() int {
	if len(s.Storage) < 1 {
		panic("stack underflow")
	}

	last := len(s.Storage) - 1
	popped := s.Storage[last]
	s.Storage = s.Storage[:last]

	return popped
}

func main() {
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
}
