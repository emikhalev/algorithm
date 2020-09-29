// https://leetcode.com/problems/rotting-oranges/
package main

func orangesRotting(grid [][]int) int {

	minutes := 0
	for {
		noFresh, changed := turn(grid)
		if noFresh && !changed {
			return minutes
		}
		if !noFresh && !changed {
			return -1
		}
		minutes++
	}

	return -1
}

func turn(grid [][]int) (noFresh, changed bool) {
	noFresh = true
	changed = false

	gridCpy := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		gridCpy[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			gridCpy[i][j] = grid[i][j]
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				noFresh = false
				if anyRotten(gridCpy, i, j) {
					grid[i][j] = 2
					changed = true
				}
			}
		}
	}

	return noFresh, changed
}

func anyRotten(grid [][]int, r, c int) bool {
	rc := len(grid)
	if rc == 0 {
		return false
	}
	cc := len(grid[0])

	// up
	if r-1 >= 0 && grid[r-1][c] == 2 {
		return true
	}

	// down
	if r+1 < rc && grid[r+1][c] == 2 {
		return true
	}

	// left
	if c-1 >= 0 && grid[r][c-1] == 2 {
		return true
	}

	// right
	if c+1 < cc && grid[r][c+1] == 2 {
		return true
	}

	return false
}
