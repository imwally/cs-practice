package doublylinkedlist

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (l *DoublyLinkedList) Append(v interface{}) {
	node := &Node{v, nil, nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}

	node.Prev = l.Tail
	l.Tail.Next = node
	l.Tail = node
}

func (l *DoublyLinkedList) Remove(v interface{}) {
	if l.Head.Value == v {
		l.Head = l.Head.Next
		return
	}

	last := l.Head
	for current := l.Head; current != nil; current = current.Next {
		if current.Value == v {
			last.Next = current.Next
			if current == l.Tail {
				l.Tail = last
			}
			return
		}
		last = current
	}
}

func (l *DoublyLinkedList) Size() int {
	i := 0
	for current := l.Head; current != nil; current = current.Next {
		i++
	}

	return i
}

func (l *DoublyLinkedList) String() string {
	var list string
	for current := l.Head; current != nil; current = current.Next {
		list += fmt.Sprintf("%v", current.Value)
		if current.Next != nil {
			list += fmt.Sprintf(" -> ")
		}
	}

	return list
}

func (l *DoublyLinkedList) Reverse() {
	var last *Node
	for current := l.Head; current != nil; current = current.Prev {
		current.Prev = current.Next
		current.Next = last
		last = current
	}

	oldHead := l.Head
	l.Head = l.Tail
	l.Tail = oldHead
}

func PrintReversed(l *DoublyLinkedList) {
	for current := l.Tail; current != nil; current = current.Prev {
		fmt.Printf("%v", current.Value)
		if current.Prev != nil {
			fmt.Printf(" ->")
		}
	}
	fmt.Println()
}
