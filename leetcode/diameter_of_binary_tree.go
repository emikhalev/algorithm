// https://leetcode.com/problems/diameter-of-binary-tree/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	paths := make([]int, 0, 16)
	DFS(root, &paths)
	return getMax(paths)
}

// DFS traversal return max path for node
func DFS(n *TreeNode, paths *[]int) int {
	if n == nil {
		return 0
	}

	maxL := DFS(n.Left, paths)
	maxR := DFS(n.Right, paths)

	*paths = append(*paths, maxL+maxR)

	return max(maxL, maxR) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func getMax(paths []int) int {
	max := -1 << 32
	for i := range paths {
		if paths[i] > max {
			max = paths[i]
		}
	}
	return max
}
