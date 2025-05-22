/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
package lc

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	var backtrack func(i, j, k int) bool
	backtrack = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}

		visited[i][j] = true
		defer func() { visited[i][j] = false }()
		if newI, newJ := i+1, j; 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !visited[newI][newJ] {
			if backtrack(newI, newJ, k+1) {
				return true
			}
		}
		if newI, newJ := i-1, j; 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !visited[newI][newJ] {
			if backtrack(newI, newJ, k+1) {
				return true
			}
		}
		if newI, newJ := i, j+1; 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !visited[newI][newJ] {
			if backtrack(newI, newJ, k+1) {
				return true
			}
		}
		if newI, newJ := i, j-1; 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !visited[newI][newJ] {
			if backtrack(newI, newJ, k+1) {
				return true
			}
		}

		return false
	}

	for i, row := range board {
		for j := range row {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end
