/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package lc

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func sortedArrayToBST(nums []int) *TreeNode {
	return middle(nums, 0, len(nums)-1)
}

func middle(nums []int, left, right int) *TreeNode {
	if right < left {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = middle(nums, left, mid-1)
	root.Right = middle(nums, mid+1, right)
	return root
}

// @lc code=end
