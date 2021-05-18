package graph

import "fmt"

type Node struct {
	Value interface{}
	Edges map[*Node]int
}

type Graph struct {
	Nodes []*Node
}

func New() Graph {
	return Graph{}
}

func (n *Node) AddEdge(n2 *Node, distance int) {
	n.Edges[n2] = distance
}

func (n *Node) Distance(n2 *Node) int {
	return n.Edges[n2]
}

// Backwards compatibility
func (n *Node) AddChild(child *Node, distance int) {
	n.AddEdge(child, distance)
}

func (g *Graph) AddNode(value int) *Node {
	m := make(map[*Node]int)
	node := &Node{value, m}
	g.Nodes = append(g.Nodes, node)

	return node
}

func (g *Graph) NodeExists(n *Node) bool {
	for _, node := range g.Nodes {
		if n == node {
			return true
		}
	}

	return false
}

func (g *Graph) Adjacent(n1, n2 *Node) bool {
	_, ok := n1.Edges[n2]

	return ok
}

func (g *Graph) AddEdge(n1, n2 *Node, distance int) error {
	if !g.NodeExists(n1) {
		return fmt.Errorf("node %v does not exists", n1)
	}

	if !g.NodeExists(n2) {
		return fmt.Errorf("node %v does not exists", n2)
	}

	if g.Adjacent(n1, n2) {
		return fmt.Errorf("nodes %v and %v are already adjacent", n1, n2)
	}

	n1.AddEdge(n2, distance)
	n2.AddEdge(n1, distance)

	return nil
}
