// https://en.wikipedia.org/wiki/Adjacency_list
package adjacency_list

type Graph struct {
	adjList map[int][]int
	values  map[int]interface{}
}

func New() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
		values:  make(map[int]interface{}),
	}
}
