package main

import (
	"fmt"
	"testing"

	"../data-structures/graph"
)

func RouteExists(n1, n2 *graph.Node) bool {
	if n1 == nil {
		return false
	}

	for _, edge := range n1.Edges {
		if edge.Node == n2 {
			return true
		}

		return RouteExists(edge.Node, n2)
	}

	return false
}

func TestGraph(t *testing.T) {
	g := new(graph.Graph)

	n0 := g.AddNode(0)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n4 := g.AddNode(4)
	n5 := g.AddNode(5)
	n6 := g.AddNode(6)

	n0.AddChild(n1, 1)
	n1.AddChild(n2, 1)
	n2.AddChild(n4, 1)
	n3.AddChild(n2, 1)
	n4.AddChild(n5, 1)
	n5.AddChild(n6, 1)

	fmt.Println(RouteExists(n0, n6))
}
