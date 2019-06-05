package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
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

func (l *LinkedList) String() string {
	var list string
	for current := l.Head; current != nil; current = current.Next {
		list += fmt.Sprintf("%v ", current.Value)
		if current.Next != nil {
			list += fmt.Sprintf("-> ")
		}
	}

	return list
}

func main() {
	ll := &LinkedList{nil, nil}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	fmt.Println(ll)
}
