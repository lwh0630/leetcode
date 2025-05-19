/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
package lc

func numIslands(grid [][]byte) (ret int) {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ret++
				dfs(grid, i, j)
			}
		}
	}

	return ret
}

func dfs(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}

// @lc code=end
