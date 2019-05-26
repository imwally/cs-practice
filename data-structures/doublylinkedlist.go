package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(v interface{}) {
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

func (l *LinkedList) Size() int {
	i := 0
	for current := l.Head; current != nil; current = current.Next {
		i++
	}

	return i
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

func PrintReversed(l *LinkedList) {
	for current := l.Tail; current != nil; current = current.Prev {
		fmt.Printf("%v ", current.Value)
		if current.Prev != nil {
			fmt.Printf("-> ")
		}
	}
	fmt.Println()
}

func main() {
	ll := &LinkedList{nil, nil}
	fmt.Println("Size:", ll.Size())
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	fmt.Println(ll)
	PrintReversed(ll)
	fmt.Println("Size:", ll.Size())
}
