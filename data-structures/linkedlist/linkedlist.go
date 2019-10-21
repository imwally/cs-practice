package linkedlist

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Size() int {
	i := 0
	for current := l.Head; current != nil; current = current.Next {
		i++
	}

	return i
}

func (l *LinkedList) Append(v interface{}) {
	node := &Node{v, nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}

	l.Tail.Next = node
	l.Tail = node
}

func (l *LinkedList) Insert(at int, v interface{}) {
	if at == 0 {
		node := &Node{v, l.Head}
		l.Head = node
		return
	}

	if at >= l.Size() {
		l.Append(v)
		return
	}

	last := l.Head
	for i, current := 0, l.Head; current != nil; i, current = i+1, current.Next {
		if i == at {
			last.Next = &Node{v, last.Next}
			return
		}
		last = current
	}
}

func (l *LinkedList) Delete(at int) {
	if at == 0 {
		l.Head = l.Head.Next
		return
	}

	last := l.Head
	for i, current := 0, l.Head; current != nil; i, current = i+1, current.Next {
		if i == at {
			last.Next = current.Next
			if at == l.Size() {
				l.Tail = last
			}

			return
		}
		last = current
	}
}

func (l *LinkedList) String() string {
	var list string
	for current := l.Head; current != nil; current = current.Next {
		list += fmt.Sprintf("%v", current.Value)
		if current.Next != nil {
			list += fmt.Sprintf(" -> ")
		}
	}

	return list
}
