/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
package lc

func subsets(nums []int) [][]int {
	ans := [][]int{}
	set := []int{}

	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}

		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}

// @lc code=end
