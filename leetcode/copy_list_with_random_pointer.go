// https://leetcode.com/problems/copy-list-with-random-pointer/
package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	c := head
	if c == nil {
		return nil
	}

	to := new(Node)
	toHead := to
	var pTo *Node = nil

	for c != nil {
		to.Val = c.Val
		c = c.Next

		pTo = to
		to.Next = new(Node)
		to = to.Next
	}
	pTo.Next = nil

	rndLinks := randomLinks(head)
	toAddrs := addresses(toHead)
	c = head
	to = toHead
	i := 0
	for c != nil {
		lnk := rndLinks[i]
		if lnk != -1 {
			to.Random = toAddrs[lnk]
		}

		c = c.Next
		to = to.Next
		i++
	}

	return toHead
}

func addresses(root *Node) []*Node {
	a := make([]*Node, 0, 16)
	c := root
	i := 0
	for ; c != nil; i++ {
		a = append(a, c)
		c = c.Next
	}
	return a
}

func randomLinks(root *Node) []int {
	addr2Idx := make(map[*Node]int)
	c := root
	i := 0
	for c != nil {
		addr2Idx[c] = i
		i++
		c = c.Next
	}

	rLinks := make([]int, 0, 16)
	c = root
	for c != nil {
		idx, ok := addr2Idx[c.Random]
		if !ok {
			idx = -1
		}
		rLinks = append(rLinks, idx)
		c = c.Next
	}
	return rLinks
}
