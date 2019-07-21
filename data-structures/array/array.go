package main

import "fmt"

type Array struct {
	Data []interface{}
	Len  int
	Cap  int
}

func (a *Array) grow() {
	newCap := a.Cap * 2
	newData := make([]interface{}, newCap)

	copy(newData, a.Data)

	a.Data = newData
	a.Cap = newCap
}

func (a *Array) shrink() {
	newCap := int(a.Cap / 2)
	newData := make([]interface{}, newCap)

	copy(newData, a.Data)

	a.Data = newData
	a.Cap = newCap
}

func (a *Array) Append(v interface{}) {
	i := a.Len
	if i == a.Cap {
		a.grow()
	}

	a.Data[i] = v
	a.Len++
}

func (a *Array) Pop() interface{} {
	a.Len--

	if a.Len < int(a.Cap/2) {
		a.shrink()
	}

	return a.Data[a.Len]
}

func (a *Array) Size() int {
	return a.Len
}

func (a *Array) String() string {
	s := "["
	for i := 0; i < a.Len; i++ {
		s += fmt.Sprintf("%v", a.Data[i])
		if i < a.Len-1 {
			s += ", "
		}
	}
	s += "]"

	return s
}

func New() *Array {
	data := make([]interface{}, 10)

	return &Array{data, 0, 10}
}

func main() {
	a := New()
	for i := 1; i <= 30; i++ {
		a.Append(i)
	}

	fmt.Println(a)
	fmt.Println(a.Size())
	fmt.Println(a.Cap)

	for i := a.Size(); i > 0; i-- {
		fmt.Println(a.Pop())
		fmt.Println(a.Data)
	}
}
