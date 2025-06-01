/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
package lc

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32 // 初始化最小价格为最大整数
	maxProfit := 0            // 初始化最大利润为0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // 更新历史最低价
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit // 更新最大利润
		}
	}
	return maxProfit
}

// @lc code=end
