// https://en.wikipedia.org/wiki/Depth-first_search
package algorithms

import (
	"github.com/emikhalev/algorithm/data_structures/graph/adjacency_list"
)

func DFS(graph *adjacency_list.Graph, startVertexIdx int,
	fnForwardCallback, fnBackCallback func(nodeIdx int)) {

	visited := make(map[int]struct{})
	dfs(graph, visited, startVertexIdx, fnForwardCallback, fnBackCallback)
}

func dfs(graph *adjacency_list.Graph, visited map[int]struct{},
	v int, fnForward, fnBack func(nodeIdx int)) {

	fnForward(v)
	visited[v] = struct{}{}
	adjEdges := graph.Neighbors(v)
	for _, w := range adjEdges {
		if _, ok := visited[w]; !ok {
			dfs(graph, visited, w, fnForward, fnBack)
		}
	}

	fnBack(v)
}
