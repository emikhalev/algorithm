package adjacency_list

func (g *Graph) Adjacent(x, y int) bool {
	if nodeX, ok := g.adjList[x]; ok {
		for _, v := range nodeX {
			if v == y {
				return true
			}
		}
	}
	return false
}

func (g *Graph) Neighbors(x int) []int {
	if nodeX, ok := g.adjList[x]; ok {
		return nodeX
	}
	return []int{}
}

func (g *Graph) AddVertex(x int) {
	if _, ok := g.adjList[x]; !ok {
		g.adjList[x] = make([]int, 0, 4)
		g.values[x] = nil
	}
}

func (g *Graph) RemoveVertex(x int) {
	if _, ok := g.adjList[x]; ok {
		delete(g.adjList, x)
		delete(g.values, x)
	}
}

func (g *Graph) AddEdge(x, y int) {
	if g.Adjacent(x, y) {
		return
	}
	nodeX, _ := g.adjList[x]
	g.adjList[x] = append(nodeX, y)
}

func (g *Graph) RemoveEdge(x, y int) {
	if nodeX, ok := g.adjList[x]; ok {
		for idx, _ := range nodeX {
			if nodeX[idx] == y {
				// Fast delete element from slice for golang
				// Move last element to the position of element, that would be deleted
				// And truncates slice (remove last element)
				nodeX[idx] = nodeX[len(nodeX)-1]
				g.adjList[x] = nodeX[0 : len(nodeX)-1]
			}
		}
	}
}

func (g *Graph) GetVertexValue(x int) (v interface{}, ok bool) {
	v, ok = g.values[x]
	return v, ok
}

func (g *Graph) SetVertexValue(x int, v interface{}) (ok bool) {
	if _, ok := g.adjList[x]; !ok {
		return false
	}
	g.values[x] = v
	return true
}
