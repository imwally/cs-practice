package graph

import "testing"

func TestGraphAddNode(t *testing.T) {
	g := &Graph{}

	for i := 1; i <= 10; i++ {
		g.AddNode(i)
	}

	for i, node := range g.Nodes {
		if node.Value != i+1 {
			t.Error("Graph.AddNode failed")
		}
	}
}
