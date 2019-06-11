package heap

import "fmt"

type Heap struct {
	Data      []int
	NextIndex int
}

func swap(data []int, a, b int) {
	data[a], data[b] = data[b], data[a]
}

func New() *Heap {
	data := make([]int, 10)

	return &Heap{data, 0}
}

func (h *Heap) HeapifyUp() {
	child := h.NextIndex
	for child > 0 {
		parent := int((child - 1) / 2)
		if h.Data[child] < h.Data[parent] {
			swap(h.Data, parent, child)
		}
		child = parent
	}
}

func (h *Heap) HeapifyDown() {
	p := 0

	for p < h.NextIndex {
		l, r := 2*p+1, 2*p+2
		s, si := h.Data[p], p

		if l < h.NextIndex && h.Data[l] < s {
			s, si = h.Data[l], l
		}
		if r < h.NextIndex && h.Data[r] < s {
			s, si = h.Data[r], r
		}

		if p == si {
			return
		}

		swap(h.Data, p, si)
		p = si
	}
}

func (h *Heap) Insert(v int) {
	if h.NextIndex == len(h.Data) {
		newData := make([]int, len(h.Data)*2)
		copy(newData, h.Data)
		h.Data = newData
	}

	h.Data[h.NextIndex] = v
	h.HeapifyUp()
	h.NextIndex++
}

func (h *Heap) Remove() int {
	h.NextIndex--
	root := h.Data[0]
	swap(h.Data, 0, h.NextIndex)
	h.HeapifyDown()

	return root
}

func (h *Heap) String() string {
	s := "["
	for i := 0; i < h.NextIndex; i++ {
		s += fmt.Sprintf("%d", h.Data[i])
		if i+1 < h.NextIndex {
			s += fmt.Sprintf(", ")
		}
	}
	s += "]"

	return s
}
