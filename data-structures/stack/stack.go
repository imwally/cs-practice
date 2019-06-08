package stack

import "fmt"

type Stack struct {
	Storage []int
}

func (s *Stack) Size() int {
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
	if s.Size() < 1 {
		panic("stack underflow")
	}

	last := len(s.Storage) - 1
	return s.Storage[last]
}

func (s Stack) String() string {
	if s.Size() < 1 {
		return ""
	}

	var str string
	for i := s.Size() - 1; i >= 0; i-- {
		str += fmt.Sprintf("%d, ", s.Storage[i])
	}

	// Everything but the last comma and space
	return str[:len(str)-2]
}
