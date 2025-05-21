/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
package lc

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	combinations = []string{}
	if len(digits) == 0 {
		return combinations
	}

	var dfs func(digits string, start int, combination string)
	dfs = func(digits string, start int, combination string) {
		if start == len(digits) {
			combinations = append(combinations, combination)
			return
		}
		digit := string(digits[start])
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			dfs(digits, start+1, combination+string(letters[i]))
		}
	}
	dfs(digits, 0, "")
	return combinations
}

// @lc code=end
