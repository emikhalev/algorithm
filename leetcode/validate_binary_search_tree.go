// https://leetcode.com/problems/validate-binary-search-tree/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := -2 << 31
	r := 2 << 31

	return isValidBST2(root, l, r)
}

func isValidBST2(root *TreeNode, l, r int) bool {
	if root == nil {
		return true
	}
	v := root.Val

	lV := l
	if root.Left != nil {
		lV = root.Left.Val
	}

	rV := r
	if root.Right != nil {
		rV = root.Right.Val
	}

	isOk := lV < v && v < rV &&
		isValidBST2(root.Left, l, v) && isValidBST2(root.Right, v, r)

	return isOk
}
