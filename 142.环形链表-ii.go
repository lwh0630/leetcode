/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package lc

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next
	for slow != fast {
		if slow == nil || fast == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// @lc code=end
