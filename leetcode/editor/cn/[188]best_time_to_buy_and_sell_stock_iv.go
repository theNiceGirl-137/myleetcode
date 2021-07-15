//给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。 
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。 
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
//
// 
//
// 示例 1： 
//
// 
//输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。 
//
// 示例 2： 
//
// 
//输入：k = 2, prices = [3,2,6,5,0,3]
//输出：7
//解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 
//。 
//
// 
//
// 提示： 
//
// 
// 0 <= k <= 100 
// 0 <= prices.length <= 1000 
// 0 <= prices[i] <= 1000 
// 
// Related Topics 数组 动态规划 
// 👍 528 👎 0


package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {
	days := len(prices)
	if days < 2 {
		return 0
	}
	if k >= days {
		return maxProfitUnlimited(prices)
	}
	buy := make([]int, k+1)
	for i := range buy {
		buy[i] = math.MinInt32
	}
	sell := make([]int, k+1)
	for i := 0; i < days; i++ {
		// 表示第几次进行买卖
		// 我们发现一旦可以赚钱就进行买卖
		// 对于每天的股票价格，表示第 j 次买入或者卖出的最大收益
		// 一次操作包括买入和卖出两种行为，对于进行买入操作时，价格尽可能低才能保证其后能卖出跟高的价格
		// 进行卖出操作时，价格尽可能高才能获得更好的收益
		for j := 1; j <= k; j++ {
			buy[j] = max188(buy[j], sell[j-1]-prices[i])
			sell[j] = max188(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k]
}

func maxProfitUnlimited(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func max188(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
