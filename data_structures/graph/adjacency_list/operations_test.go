package adjacency_list

import (
	"testing"
)

var (
	graph *Graph

	//   0  1  2  3  4 comment is here because of go fmt ...
	testGraphAdjMaxtrix = [][]int{
		{0, 1, 1, 0, 0}, // 0
		{0, 0, 0, 0, 0}, // 1
		{0, 0, 0, 0, 0}, // 2
		{0, 0, 1, 0, 0}, // 3
		{0, 0, 0, 0, 1}, // 4
	}
)

func TestMain(m *testing.M) {
	initGraph()
	m.Run()
}

func initGraph() {
	graph = New()
	for idx, _ := range testGraphAdjMaxtrix {
		graph.AddVertex(idx)
	}
	for x, _ := range testGraphAdjMaxtrix {
		for y, _ := range testGraphAdjMaxtrix {
			if testGraphAdjMaxtrix[x][y] == 1 {
				graph.AddEdge(x, y)
			}
		}
	}
}

func TestGraph_Adjacent(t *testing.T) {
	for x, _ := range testGraphAdjMaxtrix {
		for y, _ := range testGraphAdjMaxtrix {
			got := graph.Adjacent(x, y)
			exp := testGraphAdjMaxtrix[x][y] == 1
			if got != exp {
				t.Errorf("Test adjacent (x, y: %d, %d), got %v, expected %v",
					x, y, got, exp)
			}
		}
	}
}
