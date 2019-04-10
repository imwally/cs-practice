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

func (s *Stack) Peek() int {
	if len(s.Storage) < 1 {
		panic("stack underflow")
	}

	last := len(s.Storage) - 1
	return s.Storage[last]
}

func main() {
	s := new(Stack)
	fmt.Println("Push: 1")
	s.Push(1)
	fmt.Println("Push: 2")
	s.Push(2)
	fmt.Println("Push: 3")
	s.Push(3)

	fmt.Println(s)
	fmt.Println("Pop:", s.Pop())
	fmt.Println(s)
	fmt.Println("Peek:", s.Peek())
	fmt.Println("Pop:", s.Pop())
	fmt.Println(s)
	fmt.Println("Pop:", s.Pop())
	fmt.Println("Peek:", s.Peek())
}
