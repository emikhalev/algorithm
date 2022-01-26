// https://leetcode.com/problems/game-of-life/solution/
package simple_solution

func gameOfLife(board [][]int) {
	bc := copyBoard(board)
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			nCnt := neighbors(bc, r, c)
			board[r][c] = newState(bc[r][c], nCnt)
		}
	}
}

func newState(curState, nCnt int) int {
	if curState == 0 {
		if nCnt == 3 {
			return 1
		}
		return 0
	} else {
		if nCnt < 2 {
			return 0
		}
		if nCnt == 2 || nCnt == 3 {
			return 1
		}
		if nCnt > 3 {
			return 0
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
			if i >= 0 && i < rCnt && j >= 0 && j < cCnt && board[i][j] != 0 {
				cnt++
			}
		}
	}
	return cnt
}

func copyBoard(board [][]int) [][]int {
	newBoard := make([][]int, len(board))
	for r := 0; r < len(board); r++ {
		newBoard[r] = make([]int, len(board[r]))
		copy(newBoard[r], board[r])
	}
	return newBoard
}
