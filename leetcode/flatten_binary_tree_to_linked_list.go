// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	l := root.Left // Linked List current node
	r := root.Right
	llc := root
	if nllc := process(llc, l); nllc != nil {
		llc = nllc
	}
	process(llc, r)
}

func process(llc, c *TreeNode) *TreeNode {
	if c == nil {
		return nil
	}

	llc.Right = c
	llc.Left = nil
	llc = c

	l := c.Left
	r := c.Right
	if nllc := process(llc, l); nllc != nil {
		llc = nllc
	}
	if nllc := process(llc, r); nllc != nil {
		llc = nllc
	}

	return llc
}
