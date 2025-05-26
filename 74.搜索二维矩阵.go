/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
package lc

type pose struct {
	x int
	y int
}

func (p1 pose) LessOrEqual(p2 pose) bool {
	if p1.x < p2.x {
		return true
	} else if p1.x == p2.x && p1.y <= p2.y {
		return true
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	begin := pose{0, 0}
	end := pose{m - 1, n - 1}

	for begin.LessOrEqual(end) {
		l := (begin.x*n + begin.y + end.x*n + end.y) / 2
		mid := pose{l / n, l % n}
		if target == matrix[mid.x][mid.y] {
			return true
		} else if target < matrix[mid.x][mid.y] {
			l = l - 1
			end = pose{l / n, l % n}
		} else {
			l = l + 1
			begin = pose{l / n, l % n}
		}
	}

	return false
}

// @lc code=end
