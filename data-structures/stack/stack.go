package stack

import "fmt"

type Stack struct {
	Storage []int
}

func (s *Stack) Length() int {
	return len(s.Storage)
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
	if s.Length() < 1 {
		panic("stack underflow")
	}

	last := len(s.Storage) - 1
	return s.Storage[last]
}

func (s Stack) String() string {
	if s.Length() < 1 {
		return ""
	}

	var str string
	for i := s.Length() - 1; i >= 0; i-- {
		str += fmt.Sprintf("%d, ", s.Storage[i])
	}

	// Everything but the last comma and space
	return str[:len(str)-2]
}
