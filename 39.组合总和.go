/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
package lc

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return ret
	}

	var dfs func(index int, sum int)
	tmp := []int{}
	dfs = func(index int, sum int) {
		if index == len(candidates) {
			return
		}
		if sum == target {
			ret = append(ret, append([]int{}, tmp...))
			return
		}

		dfs(index+1, sum)
		if sum <= target {
			tmp = append(tmp, candidates[index])
			dfs(index, sum+candidates[index])
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(0, 0)

	return ret
}

// @lc code=end
