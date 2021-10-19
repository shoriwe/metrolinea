package test

import (
	"github.com/shoriwe/metrolinea/internal/data/graph"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := graph.NewGraph()

	g.AddNodes(
		[]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
	)

	g.AddRoutes(
		map[string]graph.Route{
			"1": {
				Source: "A",
				Length: 10,
				Target: "B",
			},
			"2": {
				Source: "A",
				Length: 2,
				Target: "C",
			},
			"3": {
				Source: "A",
				Length: 5,
				Target: "D",
			},
			"4": {
				Source: "D",
				Length: 50,
				Target: "C",
			},
			"5": {
				Source: "B",
				Length: 2,
				Target: "D",
			},
			"6": {
				Source: "F",
				Length: 20,
				Target: "J",
			},
			"7": {
				Source: "D",
				Length: 4,
				Target: "F",
			},
			"8": {
				Source: "G",
				Length: 5,
				Target: "F",
			},
			"9": {
				Source: "G",
				Length: 10,
				Target: "B",
			},
			"10": {
				Source: "G",
				Length: 10,
				Target: "J",
			},
			"11": {
				Source: "F",
				Length: 5,
				Target: "G",
			},
		},
	)

	route, nodeNotFound := g.Dijkstra("A", "J")
	if nodeNotFound != "" {
		t.Fatal(nodeNotFound)
	}
	reference := []string{"J", "G", "F", "D", "A"}
	if len(route) != len(reference) {
		t.Fatal(route)
	}
	for index, value := range reference {
		if route[index] != value {
			t.Fatal(route)
		}
	}
}
