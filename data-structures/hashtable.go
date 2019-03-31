package main

import "fmt"

const ARRAYSIZE int = 2096

type Record struct {
	Key   int
	Value string
	Next  *Record
}

type Table struct {
	Storage [ARRAYSIZE]*Record
}

func hash(s string) int {
	sum := 0
	for i, c := range s {
		sum += int(c)
		sum *= i
	}

	return sum % ARRAYSIZE
}

func (t *Table) Insert(s string) {
	h := hash(s)
	r := Record{h, s, nil}

	if t.Storage[h] == nil {
		t.Storage[h] = &r
	} else {
		var node *Record
		for node = t.Storage[h]; node.Next != nil; node = node.Next {
		}
		node.Next = &r
	}
}

func main() {
	t := new(Table)

	t.Insert("abcdef")
	t.Insert("abcdef")
	t.Insert("abcdef")

	key := hash("abcdef")
	fmt.Println(t.Storage[key])
	fmt.Println(t.Storage[key].Next)
	fmt.Println(t.Storage[key].Next.Next)
}
