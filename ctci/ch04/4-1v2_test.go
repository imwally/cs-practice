package main

import (
	"testing"

	"../data-structures/graph"
)

func in(n *graph.Node, v []*graph.Node) bool {
	for _, node := range v {
		if n == node {
			return true
		}
	}

	return false
}

func RouteExists(n1, n2 *graph.Node, visited []*graph.Node) bool {
	if n1 == nil {
		return false
	}

	visited = append(visited, n1)

	for _, edge := range n1.Edges {
		if edge.Node == n2 {
			return true
		}

		if !in(edge.Node, visited) {
			return RouteExists(edge.Node, n2, visited)
		}

	}

	return false
}

func TestRouteExists(t *testing.T) {
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
	n2.AddChild(n0, 1)
	n2.AddChild(n4, 1)
	n3.AddChild(n2, 1)
	n4.AddChild(n5, 1)
	n5.AddChild(n6, 1)

	var visited []*graph.Node
	got, expected := RouteExists(n0, n6, visited), true
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = RouteExists(n6, n0, visited), false
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
