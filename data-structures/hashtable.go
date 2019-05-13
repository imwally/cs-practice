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
	for _, c := range s {
		sum += int(c) * 31
	}

	return sum % ARRAYSIZE
}

func (t *Table) Insert(key, s string) {
	h := hash(key)
	record := &Record{key, s, nil}

	for r := t.Storage[h]; r != nil; r = r.Next {
		if r.Next == nil {
			r.Next = record
			return
		}
	}

	t.Storage[h] = record
}

func (t *Table) Get(key string) string {
	h := hash(key)

	for r := t.Storage[h]; r != nil; r = r.Next {
		if r.Key == key {
			return r.Value
		}
	}

	return ""
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
