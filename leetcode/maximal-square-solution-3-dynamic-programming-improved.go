// https://leetcode.com/problems/maximal-square/solution/
package main

func maximalSquare(matrix [][]byte) int {
	rc := len(matrix)
	if rc == 0 {
		return 0
	}
	cc := len(matrix[0])

	dp := make([]int, cc+1)

	maxSize := 0
	for r := 1; r <= rc; r++ {
		prev := 0
		for c := 1; c <= cc; c++ {
			temp := dp[c]
			if matrix[r-1][c-1] == '1' {
				dp[c] = min(dp[c-1], min(prev, dp[c])) + 1
				maxSize = max(maxSize, dp[c])
			} else {
				dp[c] = 0
			}
			prev = temp
		}
	}

	return maxSize * maxSize
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
