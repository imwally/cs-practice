package main

import "fmt"

const ARRAYSIZE int = 16

type Record struct {
	Key   string
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

func (t *Table) Insert(key, s string) {
	h := hash(key)
	record := &Record{key, s, nil}

	if t.Storage[h] == nil {
		t.Storage[h] = record
		return
	}

	var r *Record
	for r = t.Storage[h]; r.Next != nil; r = r.Next {
	}

	r.Next = record
}

func (t *Table) Get(key string) string {
	h := hash(key)

	if t.Storage[h] == nil {
		return ""
	}

	var value string
	for r := t.Storage[h]; r != nil; r = r.Next {
		if r.Key == key {
			value = r.Value
		}
	}

	return value
}

func main() {
	t := new(Table)

	t.Insert("wally", "Wally Jones")
	t.Insert("test", "This is a test")
	t.Insert("hello", "Hello, World!")
	t.Insert("tea", "Matcha")
	t.Insert("coffee", "Black")

	fmt.Println(t.Get("wally"))
	fmt.Println(t.Get("test"))
	fmt.Println(t.Get("hello"))
	fmt.Println(t.Get("tea"))
	fmt.Println(t.Get("coffee"))
	fmt.Println(t.Get("null"))

	fmt.Println(t.Storage)
}
