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

func TestNodeAddChild(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}

	n1.AddChild(n2, 1)
	if n1.Edges[0].Node != n2 {
		t.Error("Node.AddChild failed")
	}
	if n1.Edges[0].Distance != 1 {
		t.Error("Node.AddChild failed")
	}

	n1.AddChild(n3, 10)
	if n1.Edges[1].Node != n3 {
		t.Error("Node.AddChild failed")
	}
	if n1.Edges[1].Distance != 10 {
		t.Error("Node.AddChild failed")
	}

}
