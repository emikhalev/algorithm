// https://leetcode.com/problems/invert-binary-tree/
package invert_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	l := root.Left
	r := root.Right

	invertTree(l)
	invertTree(r)

	root.Left = r
	root.Right = l

	return root
}
