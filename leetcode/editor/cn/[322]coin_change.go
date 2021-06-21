//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回
// -1。 
//
// 你可以认为每种硬币的数量是无限的。 
//
// 
//
// 示例 1： 
//
// 
//输入：coins = [1, 2, 5], amount = 11
//输出：3 
//解释：11 = 5 + 5 + 1 
//
// 示例 2： 
//
// 
//输入：coins = [2], amount = 3
//输出：-1 
//
// 示例 3： 
//
// 
//输入：coins = [1], amount = 0
//输出：0
// 
//
// 示例 4： 
//
// 
//输入：coins = [1], amount = 1
//输出：1
// 
//
// 示例 5： 
//
// 
//输入：coins = [1], amount = 2
//输出：2
// 
//
// 
//
// 提示： 
//
// 
// 1 <= coins.length <= 12 
// 1 <= coins[i] <= 231 - 1 
// 0 <= amount <= 104 
// 
// Related Topics 动态规划 
// 👍 1332 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	// 每个硬币使用次数不限，本质上是完全背包问题
	if len(coins) < 1 {
		return -1
	}
	dp := make([]int, amount+1)
	// 初始化为最大值，如果初始化为最小值不会得到结果
	for i := range dp {
		dp[i] = amount+2
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				// 如果能使用硬币组成，选择数量最小的方式
				dp[i] = min322(dp[i], dp[i-coin]+1)
			}
		}
	}
	// 说明没有合适的组合
	if dp[amount] == amount+2 {
		return -1
	}
	return dp[amount]
}

func min322(i, j int) int {
	if i < j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
