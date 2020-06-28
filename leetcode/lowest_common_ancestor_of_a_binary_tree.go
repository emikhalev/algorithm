// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
package main

import (
	"fmt"
)

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP, pathQ := dfs(root, p, q, []*TreeNode{})

	//    printPath(pathP)
	//    printPath(pathQ)

	n := min(len(pathP), len(pathQ))
	lastEqIdx := -1
	for i := 0; i < n; i++ {
		if pathQ[i] != pathP[i] {
			return pathQ[lastEqIdx]
		}
		lastEqIdx = i
	}

	if lastEqIdx != -1 {
		return pathQ[lastEqIdx]
	}
	return nil
}

func printPath(path []*TreeNode) {
	for i := 0; i < len(path); i++ {
		fmt.Printf("%v->", path[i].Val)
	}
	fmt.Println()
}

func dfs(node *TreeNode, p, q *TreeNode, path []*TreeNode) (pathP, pathQ []*TreeNode) {
	if node == p {
		pathP = make([]*TreeNode, len(path)+1)
		copy(pathP, path)
		pathP[len(pathP)-1] = node
	}
	if node == q {
		pathQ = make([]*TreeNode, len(path)+1)
		copy(pathQ, path)
		pathQ[len(pathQ)-1] = node
	}

	if node == nil {
		return pathP, pathQ
	}

	path = append(path, node)
	pathPLeft, pathQLeft := dfs(node.Left, p, q, path)
	pathPRight, pathQRight := dfs(node.Right, p, q, path)

	return getNotEmptyDef(pathPLeft, pathPRight, pathP), getNotEmptyDef(pathQLeft, pathQRight, pathQ)
}

func getNotEmptyDef(path1, path2, def []*TreeNode) []*TreeNode {
	if len(path1) > 0 {
		return path1
	}
	if len(path2) > 0 {
		return path2
	}
	return def
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
