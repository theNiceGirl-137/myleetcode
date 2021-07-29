//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。 
//
// 示例 1: 
//
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。 
//
// 示例 2: 
//
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。 
//
// 说明: 你可以假设 n 不小于 2 且不大于 58。 
// Related Topics 数学 动态规划 
// 👍 556 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	// dp 数组表示组成 i 的数的子数的最大乘积
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			// 最大的拆分成的子数和取决于目前的 dp[i] 和 当前拆分成的两个数的积
			// 或者把 i-j 的拆分最大乘积，以 j 的变化最为迭代，不需要对 j 拆分
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]
}
//leetcode submit region end(Prohibit modification and deletion)
