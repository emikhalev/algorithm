// https://leetcode.com/problems/maximal-square/solution/
package main

func maximalSquare(matrix [][]byte) int {
	rc := len(matrix)
	if rc == 0 {
		return 0
	}
	cc := len(matrix[0])

	dp := make([][]int, rc+1)
	for r := 0; r < rc+1; r++ {
		dp[r] = make([]int, cc+1)
	}

	maxSize := 0
	for r := 1; r <= rc; r++ {
		for c := 1; c <= cc; c++ {
			if matrix[r-1][c-1] == '1' {
				dp[r][c] = min(dp[r-1][c-1], min(dp[r-1][c], dp[r][c-1])) + 1
				maxSize = max(maxSize, dp[r][c])
			}
		}
	}

	return maxSize * maxSize
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
