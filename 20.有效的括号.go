/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package lc

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	leftBrackets := map[byte]bool{'(': true, '[': true, '{': true}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if v, ok := pairs[s[i]]; ok { // 右括号逻辑
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1] // 弹出匹配的左括号
		} else if leftBrackets[s[i]] { // 合法左括号入栈
			stack = append(stack, s[i])
		} else { // 非法字符
			return false
		}
	}
	return len(stack) == 0
}

// @lc code=end
