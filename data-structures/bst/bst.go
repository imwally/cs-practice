package bst

import queue "../queue"

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
		if current.Value == v {
			return current
		}
		if v < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}

func preOrder(n *Node, s []interface{}) []interface{} {
	if n == nil {
		return nil
	}

	s = append(s, n.Value)

	if n.Left != nil {
		s = preOrder(n.Left, s)
	}

	if n.Right != nil {
		s = preOrder(n.Right, s)
	}

	return s
}

func PreOrder(n *Node) []interface{} {
	var s []interface{}
	return preOrder(n, s)
}

func inOrder(n *Node, s []interface{}) []interface{} {
	if n == nil {
		return nil
	}

	if n.Left != nil {
		s = inOrder(n.Left, s)
	}

	s = append(s, n.Value)

	if n.Right != nil {
		s = inOrder(n.Right, s)
	}

	return s
}

func InOrder(n *Node) []interface{} {
	var s []interface{}
	return inOrder(n, s)
}

func postOrder(n *Node, s []interface{}) []interface{} {
	if n == nil {
		return nil
	}

	if n.Left != nil {
		s = postOrder(n.Left, s)
	}

	if n.Right != nil {
		s = postOrder(n.Right, s)
	}

	s = append(s, n.Value)

	return s
}

func PostOrder(n *Node) []interface{} {
	var s []interface{}
	return postOrder(n, s)
}

func outOrder(n *Node, s []interface{}) []interface{} {
	if n == nil {
		return nil
	}

	if n.Right != nil {
		s = outOrder(n.Right, s)
	}

	s = append(s, n.Value)

	if n.Left != nil {
		s = outOrder(n.Left, s)
	}

	return s
}

func OutOrder(n *Node) []interface{} {
	var s []interface{}
	return outOrder(n, s)
}

func BFSOrder(n *Node) []interface{} {
	var s []interface{}

	q := queue.New()
	q.Enqueue(n)

	for q.Length() > 0 {
		node := q.Dequeue()

		s = append(s, node.(*Node).Value)

		if node.(*Node).Left != nil {
			q.Enqueue(node.(*Node).Left)
		}
		if node.(*Node).Right != nil {
			q.Enqueue(node.(*Node).Right)
		}
	}

	return s
}
