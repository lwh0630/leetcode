/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package lc

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end
