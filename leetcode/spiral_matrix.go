// https://leetcode.com/problems/spiral-matrix/
package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	n := len(matrix[0]) // cols
	m := len(matrix)    // rows

	visited := make([][]int, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, n)
	}

	ans := make([]int, m*n)
	dir := 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	c, r := 0, 0
	k := 0
	for k < m*n {
		ans[k] = matrix[r][c]
		visited[r][c] = 1
		k++

		newR, newC := r+dirs[dir][0], c+dirs[dir][1]
		if newR >= m || newR < 0 || newC >= n || newC < 0 || visited[newR][newC] != 0 {
			dir = (dir + 1) % 4
		}
		r, c = r+dirs[dir][0], c+dirs[dir][1]
	}

	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
