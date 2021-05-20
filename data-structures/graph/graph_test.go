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

func TestGraphAddEdge(t *testing.T) {
	g := New()
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)

	if err := g.AddEdge(n1, n2, 100); err != nil {
		t.Error(err)
	}
	if err := g.AddEdge(n2, n3, 200); err != nil {
		t.Error(err)
	}
	if err := g.AddEdge(n1, n3, 300); err != nil {
		t.Error(err)
	}
}

func TestGraphDistance(t *testing.T) {
	g := New()
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)

	if err := g.AddEdge(n1, n2, 100); err != nil {
		t.Error(err)
	}
	if err := g.AddEdge(n2, n3, 200); err != nil {
		t.Error(err)
	}
	if err := g.AddEdge(n1, n3, 300); err != nil {
		t.Error(err)
	}

	got, expected := n1.Distance(n2), 100
	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
	}
	got, expected = n2.Distance(n3), 200
	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
	}
	got, expected = n1.Distance(n3), 300
	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
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
