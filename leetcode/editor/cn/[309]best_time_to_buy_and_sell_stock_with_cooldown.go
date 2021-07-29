//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。 
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）: 
//
// 
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。 
// 
//
// 示例: 
//
// 输入: [1,2,3,0,2]
//输出: 3 
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出] 
// Related Topics 数组 动态规划 
// 👍 824 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit309(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	have := make([]int, n)
	notHave := make([]int, n)
	// 第 0 天购买股票
	have[0] = -prices[0]
	for i := 1; i < n; i++ {
		if i == 1 {
			// 第 1 天持有着股票，可能是昨天买的，今天休息
			// 也可能是昨天休息，今天买的
			have[i] = max(have[i-1], -prices[1])
		} else {
			// 昨天就持有股票，今天休息
			// 前天卖了股票，今天买了股票
			have[i] = max(have[i-1], notHave[i-2]-prices[i])
		}
		// 昨天也没有持有，今天休息
		// 昨天持有股票，今天卖了股票
		notHave[i] = max(notHave[i-1], have[i-1]+prices[i])
	}
	return notHave[n-1]
}
//leetcode submit region end(Prohibit modification and deletion)
