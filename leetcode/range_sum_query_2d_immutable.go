// https://leetcode.com/problems/range-sum-query-2d-immutable/
package main

type NumMatrix struct {
	rowSums [][]int
	colSums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{}
	if len(matrix) == 0 {
		return nm
	}

	rC := len(matrix)
	cC := len(matrix[0])

	nm.colSums = make([][]int, rC)
	nm.rowSums = make([][]int, rC)
	for i := 0; i < rC; i++ {
		nm.colSums[i] = make([]int, cC)
		nm.rowSums[i] = make([]int, cC)
	}

	for r := 0; r < rC; r++ {
		sum := 0
		for c := 0; c < cC; c++ {
			sum += matrix[r][c]
			nm.colSums[r][c] = sum
		}
	}

	for c := 0; c < cC; c++ {
		sum := 0
		for r := 0; r < rC; r++ {
			sum += matrix[r][c]
			nm.rowSums[r][c] = sum
		}
	}

	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	if row2-row1 < col2-col1 {
		for r := row1; r <= row2; r++ {
			v := 0
			if col1 > 0 {
				v = this.colSums[r][col1-1]
			}
			sum = sum + this.colSums[r][col2] - v
		}
	} else {
		for c := col1; c <= col2; c++ {
			v := 0
			if row1 > 0 {
				v = this.rowSums[row1-1][c]
			}
			sum = sum + this.rowSums[row2][c] - v
		}
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
