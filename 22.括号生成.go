/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
package lc

func generateParenthesis(n int) []string {
	ret := []string{}
	var backtrack func(left int, right int, curString string)
	backtrack = func(left, right int, curString string) {
		if len(curString) == 2*n {
			ret = append(ret, curString)
			return
		}

		if left < n {
			backtrack(left+1, right, curString+"(")
		}
		if right < left {
			backtrack(left, right+1, curString+")")
		}
	}
	backtrack(0, 0, "")
	return ret
}

// @lc code=end
