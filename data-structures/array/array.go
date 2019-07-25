package array

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
	if a.Len == a.Cap {
		a.grow()
	}

	a.Data[a.Len] = v
	a.Len++
}

func (a *Array) Pop() interface{} {
	if a.Len < 1 {
		panic("stack underflow")
	}

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
	var s string
	for i := 0; i < a.Len; i++ {
		s += fmt.Sprintf("%v ", a.Data[i])
	}
	s = fmt.Sprintf("[%s]", s[:len(s)-1])

	return s
}

func New() *Array {
	data := make([]interface{}, 10)

	return &Array{data, 0, 10}
}
