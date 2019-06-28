package graph

import "fmt"

type Edge struct {
	Node     *Node
	Distance int
}

type Node struct {
	Value int
	Edges []*Edge
}

type Graph struct {
	Nodes []*Node
}

func (n *Node) AddChild(child *Node, distance int) {
	n.Edges = append(n.Edges, &Edge{child, distance})
}

func (g *Graph) AddNode(value int) *Node {
	node := &Node{value, nil}
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

func (g *Graph) AddEdge(n1, n2 *Node, distance int) error {
	if !g.NodeExists(n1) {
		return fmt.Errorf("node %v does not exists", n1)
	}

	if !g.NodeExists(n2) {
		return fmt.Errorf("node %v does not exists", n2)
	}

	n1.AddChild(n2, distance)
	n2.AddChild(n1, distance)

	return nil
}
