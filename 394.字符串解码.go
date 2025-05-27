/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
package lc

func decodeString(s string) string {
	stack := [][]interface{}{} // 栈存储 [multi, last_res]
	res := ""
	multi := 0

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			multi = multi*10 + int(ch-'0')
		} else if ch == '[' {
			stack = append(stack, []interface{}{multi, res})
			res, multi = "", 0
		} else if ch == ']' {
			// 弹出栈顶元素
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curMulti := top[0].(int)
			lastRes := top[1].(string)

			// 构建当前字符串
			temp := ""
			for i := 0; i < curMulti; i++ {
				temp += res
			}
			res = lastRes + temp
		} else {
			res += string(ch)
		}
	}
	return res
}

// @lc code=end
