// https://leetcode.com/problems/binary-tree-paths/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)
	paths = getPaths(root, "", paths)
	return paths
}

func getPaths(n *TreeNode, path string, paths []string) []string {
	if n == nil {
		return paths
	}
	if path != "" {
		path += "->"
	}
	path += fmt.Sprintf("%d", n.Val)

	if n.Left == nil && n.Right == nil {
		paths = append(paths, path)
	}
	paths = getPaths(n.Left, path, paths)
	paths = getPaths(n.Right, path, paths)
	return paths
}
