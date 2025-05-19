/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	depth(root, &ans)
	return ans - 1
}

func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	L := depth(root.Left, ans)
	R := depth(root.Right, ans)

	*ans = max(*ans, L+R+1)

	return max(L, R) + 1
}

// @lc code=end
