package main

import (
	"fmt"
	"sort"

	"../../data-structures/graph"
)

func shortestDistance(g *graph.Graph, n1, n2 *graph.Node) int {
	distances := make(map[*graph.Node]int)
	shortest := make(map[*graph.Node]int)

	// Set all distances to an arbitrarly high number
	for _, n := range g.Nodes {
		distances[n] = 999999
	}

	// Set start to 0
	distances[n1] = 0

	for len(distances) > 0 {
		// Sort distances low to high
		type kv struct {
			Key   *graph.Node
			Value int
		}

		var ss []kv
		for k, v := range distances {
			ss = append(ss, kv{k, v})
		}

		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value < ss[j].Value
		})

		// Set shortest path to current
		current, distance := ss[0].Key, ss[0].Value
		shortest[current] = distance

		// Remove from map
		delete(distances, current)

		// Check all edges connected to current
		for _, edge := range current.Edges {
			// Only compute distances not seen yet
			if _, ok := distances[edge.Node]; ok {
				newDistance := distance + edge.Distance
				// Update distances map if route is shorter
				if newDistance < distances[edge.Node] {
					distances[edge.Node] = newDistance
				}
			}
		}
	}

	return shortest[n2]
}

func main() {
	g := graph.Graph{}
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n4 := g.AddNode(4)
	n5 := g.AddNode(5)
	n6 := g.AddNode(6)

	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n3, 10)
	g.AddEdge(n1, n4, 3)
	g.AddEdge(n2, n3, 4)
	g.AddEdge(n3, n4, 2)
	g.AddEdge(n3, n5, 6)
	g.AddEdge(n1, n5, 18)
	g.AddEdge(n4, n6, 20)
	g.AddEdge(n3, n6, 4)

	fmt.Println(shortestDistance(&g, n1, n2)) // 1
	fmt.Println(shortestDistance(&g, n1, n3)) // 5
	fmt.Println(shortestDistance(&g, n4, n3)) // 2
	fmt.Println(shortestDistance(&g, n1, n5)) // 11
	fmt.Println(shortestDistance(&g, n1, n6)) // 9
	fmt.Println(shortestDistance(&g, n5, n6)) // 10
}
