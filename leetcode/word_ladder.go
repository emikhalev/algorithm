// https://leetcode.com/problems/word-ladder/
package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	found := false
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == beginWord {
			found = true
			break
		}
	}
	if !found {
		wordList = append(wordList, beginWord)
	}

	adj, nodes := adjList(wordList)
	return bfs(adj, nodes, beginWord, endWord)
}

func bfs(adj map[int][]int, nodes map[string]int, beginWord, endWord string) int {
	eIdx, ok := nodes[endWord]
	if !ok {
		return 0
	}

	color := make([]int, len(nodes))
	distance := make([]int, len(nodes))
	u, ok := nodes[beginWord]
	if !ok {
		return 0
	}

	q := NewQueue()
	q.Enqueue(u)
	color[u] = 1
	distance[u] = 0
	for q.Size() > 0 {
		u = q.Dequeue()
		if eIdx == u {
			return distance[u] + 1
		}
		n, _ := adj[u]
		for i := 0; i < len(n); i++ {
			v := n[i]
			if color[v] == 0 {
				q.Enqueue(v)
				color[v] = 1
				distance[v] = distance[u] + 1
			}
		}
		color[u] = 2
	}

	return 0
}

func adjList(wordList []string) (adj map[int][]int, nodes map[string]int) {
	adj = make(map[int][]int)
	nodes = make(map[string]int)

	for i := 0; i < len(wordList); i++ {
		w1 := wordList[i]
		nodes[w1] = i
		for j := i + 1; j < len(wordList); j++ {
			w2 := wordList[j]
			if isOneDiff(w1, w2) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	return adj, nodes
}

func isOneDiff(w1, w2 string) bool {
	cnt := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

type Queue struct {
	sIn, sOut *Stack
}

func NewQueue() *Queue {
	return &Queue{
		sIn:  NewStack(),
		sOut: NewStack(),
	}
}

func (q *Queue) Size() int {
	return q.sIn.Size() + q.sOut.Size()
}

func (q *Queue) Enqueue(v int) {
	q.sIn.Push(v)
}

func (q *Queue) Dequeue() int {
	if q.sOut.Size() == 0 {
		for q.sIn.Size() > 0 {
			q.sOut.Push(q.sIn.Pop())
		}
	}
	return q.sOut.Pop()
}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
	idx := len(s.data) - 1
	v := s.data[idx]
	s.data = s.data[:idx]
	return v
}

func (s *Stack) Size() int {
	return len(s.data)
}
