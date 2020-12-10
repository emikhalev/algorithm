package algorithms

import (
	"testing"

	"github.com/emikhalev/algorithm/data_structures/graph/adjacency_list"
)

var (
	graph *adjacency_list.Graph

	/*
	        0
	      /  \
	     1	--2
	        \/ \
	        3 - 4

	*/
	//   0  1  2  3  4 comment is here because of go fmt ...
	testGraphAdjMaxtrix = [][]int{
		{0, 1, 1, 0, 0}, // 0
		{0, 0, 0, 0, 0}, // 1
		{0, 0, 1, 1, 0}, // 2
		{0, 0, 0, 0, 1}, // 3
		{0, 0, 0, 0, 0}, // 4
	}
)

func TestMain(m *testing.M) {
	initGraph()
	m.Run()
}

func initGraph() {
	graph = adjacency_list.New()
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

func TestDFS(t *testing.T) {
	// Expecting: 0(in)->1(in)->1(out)->2(in)->3(in)->4(in)->4(out)->3(out)->2(out)->0(out)
	const (
		DirIn  = 0
		DirOut = 1
	)
	exp := []struct {
		id        int
		direction int
	}{
		{0, DirIn},
		{1, DirIn},
		{1, DirOut},
		{2, DirIn},
		{3, DirIn},
		{4, DirIn},
		{4, DirOut},
		{3, DirOut},
		{2, DirOut},
		{0, DirOut},
	}

	expIdx := 0
	checkFn := func(nodeIdx int, dir int) {
		if nodeIdx != exp[expIdx].id || dir != exp[expIdx].direction {
			t.Errorf("expIdx: %d) expecting id=%d, direction=%d, got id=%d, direction=%d",
				expIdx,
				exp[expIdx].id, exp[expIdx].direction,
				nodeIdx, dir)
		}
		expIdx++
	}
	DFS(graph, 0,
		func(nodeIdx int) {
			checkFn(nodeIdx, DirIn)
		},
		func(nodeIdx int) {
			checkFn(nodeIdx, DirOut)
		},
	)
}
