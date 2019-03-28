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

func (t *Table) Insert(s string) int {
	h := hash(s)
	r := Record{h, s, nil}

	if t.Storage[h] == nil {
		t.Storage[h] = &r
		return h
	}

	return h
}

func main() {
	t := new(Table)

	key := t.Insert("abcdef")

	fmt.Println(key)
	fmt.Println(t.Storage[key])
}
