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

func (l *LinkedList) Remove(v interface{}) {
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

func (l *LinkedList) Reverse() {
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
	fmt.Println("Add: 1")
	ll.Append(1)
	fmt.Println("Add: 2")
	ll.Append(2)
	fmt.Println("Add: 3")
	ll.Append(3)
	fmt.Println("Add: 4")
	ll.Append(4)

	fmt.Println("List:", ll)
	fmt.Printf("Printing Reversed: ")
	PrintReversed(ll)
	fmt.Println("Size:", ll.Size())

	fmt.Println("Reversing list.")
	ll.Reverse()
	fmt.Println("List:", ll)

	fmt.Printf("Printing Reverse: ")
	PrintReversed(ll)

	fmt.Println("Remove: 4")
	ll.Remove(4)
	fmt.Println("List:", ll)
	fmt.Println("Head:", ll.Head.Value)
	fmt.Println("Tail:", ll.Tail.Value)
	fmt.Println("Size:", ll.Size())

	fmt.Println("Remove: 3")
	ll.Remove(3)
	fmt.Println("List:", ll)
	fmt.Println("Head:", ll.Head.Value)
	fmt.Println("Tail:", ll.Tail.Value)

	ll.Remove(1)
	fmt.Println("List:", ll)
	fmt.Println("Head:", ll.Head.Value)
	fmt.Println("Tail:", ll.Tail.Value)
}
