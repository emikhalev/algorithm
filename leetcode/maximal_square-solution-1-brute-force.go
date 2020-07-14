// https://leetcode.com/problems/maximal-square/solution/
package main

func maximalSquare(matrix [][]byte) int {
	rc := len(matrix)
	if rc == 0 {
		return 0
	}
	cc := len(matrix[0])

	maxSize := 0
	for r := 0; r < rc; r++ {
		for c := 0; c < cc; c++ {
			maxSize = max(maxSize, maxSquareAtPos(r, c, matrix))
		}
	}

	return maxSize
}

func maxSquareAtPos(rs, cs int, matrix [][]byte) int {
	rc := len(matrix)
	cc := len(matrix[0])

	maxSize := 0
	for size := 1; size <= min(rc-rs, cc-cs); size++ {
		for r := rs; r < rs+size; r++ {
			for c := cs; c < cs+size; c++ {
				if matrix[r][c] != '1' {
					return maxSize * maxSize
				}
			}
		}
		maxSize = size
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
