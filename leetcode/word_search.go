// https://leetcode.com/problems/word-search/
package main

func exist(board [][]byte, word string) bool {
	rc := len(board)
	if rc == 0 {
		return false
	}
	cc := len(board[0])

	visited := make([][]int, rc)
	for i := 0; i < rc; i++ {
		visited[i] = make([]int, cc)
	}

	for r := 0; r < rc; r++ {
		for c := 0; c < cc; c++ {
			if search(board, visited, word, r, c, rc, cc) {
				return true
			}
		}
	}

	return false
}

func search(board [][]byte, visited [][]int, word string, r, c int, rc, cc int) bool {
	if word == "" {
		return true
	}
	if r < 0 || c < 0 || r >= rc || c >= cc || visited[r][c] == 1 {
		return false
	}
	if board[r][c] != word[0] {
		return false
	}

	visited[r][c] = 1

	if search(board, visited, word[1:], r+1, c, rc, cc) {
		return true
	}
	if search(board, visited, word[1:], r-1, c, rc, cc) {
		return true
	}
	if search(board, visited, word[1:], r, c+1, rc, cc) {
		return true
	}
	if search(board, visited, word[1:], r, c-1, rc, cc) {
		return true
	}

	visited[r][c] = 0
	return false
}
