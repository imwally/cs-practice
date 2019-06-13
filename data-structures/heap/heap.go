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

// HeapifyUp starts at the last element in the array and works
// backwards from child to parent checking if the value is smaller
// than its parent. If the child is smaller than the parent, they are
// swapped. This happens each time a new element is pushed onto the
// heap to maintain order.
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

// HeapifyDown starts at the root (or first element of the data array)
// and works downwards checking the parent and both children to find
// the smallest value among the three. The smallest value will be
// swapped with the parent. This happens each time an element is
// popped off the heap to maintain order.
func (h *Heap) HeapifyDown() {
	for p := 0; p < h.NextIndex; {
		// Left and right child
		l, r := 2*p+1, 2*p+2

		// Smallest and its index
		s, si := h.Data[p], p

		// Check if left child is smaller than parent, if so
		// update smallest and its index
		if l < h.NextIndex && h.Data[l] < s {
			s, si = h.Data[l], l
		}

		// Check if right is smaller than smallest seen so far
		if r < h.NextIndex && h.Data[r] < s {
			s, si = h.Data[r], r
		}

		// If parent is already the smallest do nothing
		if p == si {
			return
		}

		swap(h.Data, p, si)
		p = si
	}
}

// Push inserts a new value into the data array and then calls
// HeapifyUp to maintain order. If the array is full the size will be
// doubled and its contents copied.
func (h *Heap) Push(v int) {
	// Resize array if full
	if h.NextIndex == len(h.Data) {
		newData := make([]int, len(h.Data)*2)
		copy(newData, h.Data)
		h.Data = newData
	}

	h.Data[h.NextIndex] = v
	h.HeapifyUp()
	h.NextIndex++
}

// Pop removes and returns the root (or first element of the data
// array) from the heap. The root and last element are swapped then
// HeapifyDown is called to maintain order.
func (h *Heap) Pop() int {
	if h.NextIndex == 0 {
		panic("stack underflow")
	}

	h.NextIndex--
	root := h.Data[0]
	swap(h.Data, 0, h.NextIndex)
	h.HeapifyDown()

	// Resize array if only half is being used
	if h.NextIndex == len(h.Data)/2 {
		newData := make([]int, len(h.Data)/2)
		copy(newData, h.Data)
		h.Data = newData
	}

	return root
}

func (h *Heap) Peek() int {
	return h.Data[0]
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
