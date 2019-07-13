package main

import "fmt"

type Array struct {
	Data []interface{}
	Len  int
	Cap  int
}

func (a *Array) Append(v interface{}) {
	i := a.Len
	a.Data[i] = v

	a.Len++
}

func New() *Array {
	data := make([]interface{}, 10)

	return &Array{data, 0, 10}
}

func main() {
	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(4)

	fmt.Println(a)
}
