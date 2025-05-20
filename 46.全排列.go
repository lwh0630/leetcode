/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
package lc

func factorial(num int) (ret int) {
	ret = 1
	for i := 2; i <= num; i++ {
		ret = ret * i
	}
	return
}

func swap(i, j *int) {
	tmp := *i
	*i = *j
	*j = tmp
}

func backtrack(res *[][]int, nums []int, start int) {
	if start == len(nums) {
		// tmp := make([]int, len(nums))
		// copy(tmp, nums)
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i := start; i < len(nums); i++ {
		swap(&nums[start], &nums[i])
		backtrack(res, nums, start+1)
		swap(&nums[start], &nums[i])
	}
}

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0, factorial(len(nums)))
	backtrack(&res, nums, 0)
	return res
}

// @lc code=end
