// https://leetcode.com/problems/game-of-life/solution/
package in_place

func gameOfLife(board [][]int) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			nCnt := neighbors(board, r, c)
			board[r][c] = newState(board[r][c], nCnt)
		}
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] <= 0 {
				board[r][c] = 0
			}
			if board[r][c] > 0 {
				board[r][c] = 1
			}
		}
	}
}

func newState(curState, nCnt int) int {
	if curState == 0 {
		if nCnt == 3 {
			return 2
		}
		return 0
	}
	if curState == 1 {
		if nCnt < 2 {
			return -1
		}
		if nCnt == 2 || nCnt == 3 {
			return 1
		}
		if nCnt > 3 {
			return -1
		}
	}
	return curState
}

func neighbors(board [][]int, r, c int) int {
	rCnt := len(board)
	cCnt := len(board[0])
	cnt := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i == r && j == c {
				continue
			}
			if i >= 0 && i < rCnt && j >= 0 && j < cCnt && abs(board[i][j]) == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
