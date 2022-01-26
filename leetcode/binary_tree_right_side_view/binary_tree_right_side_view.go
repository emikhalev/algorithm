// https://leetcode.com/problems/binary-tree-right-side-view/
package binary_tree_right_side_view

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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	tH := treeHeight(root, 1)
	visited := make([]bool, tH)
	view := make([]int, tH)

	view[0] = root.Val
	treeRightView(root.Right, visited, view, 1)
	treeRightView(root.Left, visited, view, 1)

	return view
}

func treeRightView(node *TreeNode, visited []bool, view []int, level int) {
	if node == nil {
		return
	}

	if !visited[level] {
		view[level] = node.Val
		visited[level] = true
	}

	treeRightView(node.Right, visited, view, level+1)
	treeRightView(node.Left, visited, view, level+1)
}

func treeHeight(root *TreeNode, level int) int {
	if root == nil {
		return level - 1
	}
	h := max(level, treeHeight(root.Left, level+1))
	h = max(h, treeHeight(root.Right, level+1))
	return h
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
