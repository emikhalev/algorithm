// https://leetcode.com/problems/binary-tree-maximum-path-sum/
package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	_, maxSum := maxSumForNode(root)
	return maxSum
}

const intMin = -1<<31

func maxSumForNode(node *TreeNode) (val, maxSum int) {
	if node == nil {
		return intMin, intMin
	}


	valL, maxSumL := maxSumForNode(node.Left);
	valR, maxSumR := maxSumForNode(node.Right)

	sumC := node.Val
	sumCAll := node.Val
	sumCL := node.Val
	sumCR := node.Val
	if valL != intMin {
		sumCAll += valL
		sumCL += valL
	}
	if valR != intMin {
		sumCAll += valR
		sumCR += valR
	}

	//fmt.Printf("%v\n", sumCAll)
	sumCMax := max(sumC, max(sumCL, sumCR))

	return sumCMax, max(sumCAll, max(sumCMax, max(maxSumL, maxSumR)))
}

func max(i, j int) int {
	if i>j {
		return i
	}
	return j
}
