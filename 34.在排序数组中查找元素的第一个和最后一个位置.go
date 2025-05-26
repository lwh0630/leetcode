/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package lc

import "sort"

func searchRange(nums []int, target int) []int {
	leftIndex := sort.SearchInts(nums, target)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}
	rightIndex := sort.SearchInts(nums, target+1) - 1

	return []int{leftIndex, rightIndex}
}

// @lc code=end
