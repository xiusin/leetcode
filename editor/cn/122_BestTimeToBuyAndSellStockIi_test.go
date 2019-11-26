package cn

import (
	"fmt"
	"testing"
)

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。 
//
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。 
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
//
// 示例 1: 
//
// 输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
// 
//
// 示例 2: 
//
// 输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
// 
//
// 示例 3: 
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。 
// Related Topics 贪心算法 数组

//leetcode submit region begin(Prohibit modification and deletion)
// todo 没做出来, 看了别人的解析
func maxProfitII(prices []int) int {
	// 买入价格必须得小于出售价, 所以将每次的最低买入跟最高卖出统计出来求和即可.
	// 需要图形分析, 我们总是想最大化最低谷和最高峰, 放弃中间的波峰, 所以将会损失中间的收益.
	// 比如: [1,4,2,199] , 简单考虑会1买入199卖出. 其实, 中间 1,4和2-199的收益比直接1.199的高
	// 我们可以把每次的相邻数据看作为买入和卖出
	var maxProfit int
	l := len(prices)
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1] //累积连续收益
		}
	}
	return maxProfit
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBestTimeToBuyAndSellStockIi(t *testing.T) {
	fmt.Println(maxProfitII([]int{7, 1, 5, 3, 6, 4}), 7)
	fmt.Println(maxProfitII([]int{7, 6, 4, 3, 1}), 0)
	fmt.Println(maxProfitII([]int{1, 2, 3, 4, 5}), 4)
}
