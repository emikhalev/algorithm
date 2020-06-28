// https://leetcode.com/problems/binary-search-tree-iterator/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	s *Stack
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	iter.s = NewStack()
	iter.inLeft(root)
	return iter
}

func (this *BSTIterator) inLeft(n *TreeNode) {
	if n == nil {
		return
	}
	this.s.push(n)
	this.inLeft(n.Left)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return -1
	}

	n := this.s.pop()
	if n.Right != nil {
		this.inLeft(n.Right)
	}

	return n.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.s.size() > 0
}

type Stack struct {
	data []*TreeNode
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) push(v *TreeNode) {
	s.data = append(s.data, v)
}

func (s *Stack) pop() *TreeNode {
	if len(s.data) == 0 {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) size() int {
	return len(s.data)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
