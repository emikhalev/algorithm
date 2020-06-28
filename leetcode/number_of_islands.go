// https://leetcode.com/problems/number-of-islands/
package main

func numIslands(grid [][]byte) int {
	nR := len(grid)
	if nR == 0 {
		return 0
	}

	nC := len(grid[0])
	if nC == 0 {
		return 0
	}

	cnt := 0
	for c := 0; c < nC; c++ {
		for r := 0; r < nR; r++ {
			if grid[r][c] == '1' {
				dfs(grid, c, r)
				cnt++
			}
		}
	}

	return cnt
}

func dfs(grid [][]byte, c, r int) {
	nR := len(grid)
	nC := len(grid[0])

	if c < 0 || c >= nC || r < 0 || r >= nR || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, c, r-1)
	dfs(grid, c, r+1)

	dfs(grid, c-1, r)
	dfs(grid, c+1, r)
}
