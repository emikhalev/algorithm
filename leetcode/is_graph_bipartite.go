// https://leetcode.com/problems/is-graph-bipartite/
package main

func isBipartite(graph [][]int) bool {
	res := true
	parts := make(map[int]int)
	for i := 0; i < len(graph); i++ {
		part := 1
		if v, ok := parts[i]; ok {
			part = v
		}
		if dfs(graph, parts, i, part) == false {
			return false
		}
	}
	return res
}

func dfs(graph [][]int, parts map[int]int, i int, part int) bool {
	parts[i] = part
	if part == 1 {
		part = 0
	} else {
		part = 1
	}

	for idx := 0; idx < len(graph[i]); idx++ {
		j := graph[i][idx]
		vI, ok := parts[i]
		if !ok {
			vI = 0
		}
		vJ, ok := parts[j]
		if !ok {
			vJ = 0
		}

		if vJ == 0 {
			parts[j] = part
			dfs(graph, parts, j, part)
		}

		if vI == vJ {
			return false
		}
	}

	return true
}
