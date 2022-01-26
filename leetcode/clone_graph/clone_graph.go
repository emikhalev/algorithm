// https://leetcode.com/problems/clone-graph/
package clone_graph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var firstNode *Node = nil
	if node == nil {
		return firstNode
	}
	nmap := make(map[*Node]*Node)
	createFn := func(cur *Node) {
		n := &Node{
			Val:       cur.Val,
			Neighbors: make([]*Node, 0, len(cur.Neighbors)),
		}
		if firstNode == nil {
			firstNode = n
		}
		nmap[cur] = n
	}
	linkFn := func(cur *Node, child *Node) {
		nn, _ := nmap[cur]
		cn, _ := nmap[child]
		nn.Neighbors = append(nn.Neighbors, cn)
	}

	node.BFS(createFn, linkFn)

	return firstNode
}

func (s *Node) BFS(createFn func(cur *Node), linkFn func(cur *Node, child *Node)) {
	const (
		WHITE = 0
		GRAY  = 1
		BLACK = 2
	)
	type Prop struct {
		color int
	}

	props := make(map[*Node]*Prop)
	q := NewQueue(100)

	props[s] = &Prop{color: GRAY}
	q.Enqueue(s)
	createFn(s)

	for {
		u, ok := q.Dequeue()
		if !ok {
			break
		}

		for i := 0; i < len(u.Neighbors); i++ {
			w := u.Neighbors[i]
			prop, ok := props[w]
			if !ok {
				prop = &Prop{color: WHITE}
				props[w] = prop
			}

			if prop.color == WHITE {
				createFn(w)
				prop.color = GRAY
				q.Enqueue(w)
			}

			linkFn(u, w)
		}

		p, _ := props[u]
		p.color = BLACK
	}
}

type Queue struct {
	sIn, sOut *Stack
}

func NewQueue(size int) *Queue {
	return &Queue{NewStack(size), NewStack(size)}
}

func (q *Queue) Enqueue(v *Node) bool {
	return q.sIn.Push(v)
}

func (q *Queue) Dequeue() (*Node, bool) {
	if q.sOut.Size() == 0 {
		for {
			v, ok := q.sIn.Pop()
			if !ok {
				break
			}
			q.sOut.Push(v)
		}
	}
	return q.sOut.Pop()
}

type Stack struct {
	idx  int
	data []*Node
}

func NewStack(size int) *Stack {
	return &Stack{0, make([]*Node, size)}
}

func (s *Stack) Push(v *Node) bool {
	if s.idx == len(s.data) {
		return false
	}
	s.data[s.idx] = v
	s.idx++
	return true
}

func (s *Stack) Pop() (*Node, bool) {
	if s.idx-1 >= 0 {
		s.idx--
		return s.data[s.idx], true
	}
	return nil, false
}

func (s *Stack) Size() int {
	return s.idx
}
