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

func (s *Stack) Length() int {
	return len(s.Storage)
}

func (s Stack) String() string {
	if len(s.Storage) < 1 {
		return ""
	}

	var str string
	for i := s.Length() - 1; i >= 0; i-- {
		str += fmt.Sprintf("%d, ", s.Storage[i])
	}

	// Everything but the last comma and space
	return str[:len(str)-2]
}

func main() {
	s := new(Stack)
	fmt.Println("Push: 1")
	s.Push(1)
	fmt.Println("Push: 2")
	s.Push(2)
	fmt.Println("Push: 3")
	s.Push(3)

	fmt.Println("Stack:", s)
	fmt.Println("Length:", s.Length())
	fmt.Println("Pop:", s.Pop())
	fmt.Println("Stack:", s)
	fmt.Println("Peek:", s.Peek())
	fmt.Println("Pop:", s.Pop())
	fmt.Println("Length:", s.Length())
	fmt.Println("Stack:", s)
	fmt.Println("Pop:", s.Pop())
	fmt.Println("Stack:", s)
	fmt.Println("Peek:", s.Peek())
}
