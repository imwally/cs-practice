package bst

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func New(v int) *Node {
	return &Node{v, nil, nil}
}

func (n *Node) Insert(v int) {
	newNode := New(v)
	for current := n; current != nil; {
		if v < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left

		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

func (n *Node) Find(v int) *Node {
	for current := n; current != nil; {
		if v < current.Value {
			if current.Left.Value == v {
				return current.Left
			}
			current = current.Left
		} else {
			if current.Right.Value == v {
				return current.Right
			}
			current = current.Right
		}
	}

	return nil
}
