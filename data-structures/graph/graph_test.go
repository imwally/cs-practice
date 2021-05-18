package graph

import "testing"

func TestGraphAddNode(t *testing.T) {
	g := New()

	for i := 1; i <= 10; i++ {
		g.AddNode(i)
	}

	for i, node := range g.Nodes {
		if node.Value != i+1 {
			t.Error("Graph.AddNode failed")
		}
	}
}

func TestNodeAddChild(t *testing.T) {
	g := New()
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)

	n1.AddChild(n2, 1)
	if n1.Edges[n2] != 1 {
		t.Error("Node.AddChild failed")
	}

	n1.AddChild(n3, 10)
	if n1.Edges[n3] != 10 {
		t.Error("Node.AddChild failed")
	}

}
