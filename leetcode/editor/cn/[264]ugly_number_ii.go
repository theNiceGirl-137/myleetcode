//给你一个整数 n ，请你找出并返回第 n 个 丑数 。 
//
// 丑数 就是只包含质因数 2、3 和/或 5 的正整数。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 10
//输出：12
//解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：1
//解释：1 通常被视为丑数。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 1690 
// 
// Related Topics 堆 数学 动态规划 
// 👍 624 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	q2, q3, q5 := 1, 1, 1
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = 1
			continue
		}
		a, b, c := dp[q2]*2, dp[q3]*3, dp[q5]*5
		dp[i] = min(min(a, b), c)
		if dp[i] == a {
			q2++
		}
		if dp[i] == b {
			q3++
		}
		if dp[i] == c {
			q5++
		}
 	}
 	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
