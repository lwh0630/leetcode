/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
package lc

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	var mid int
	for left <= right {
		mid = (left-right)>>1 + left
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// @lc code=end
