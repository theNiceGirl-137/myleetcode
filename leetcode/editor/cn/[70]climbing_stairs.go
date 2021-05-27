//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。 
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？ 
//
// 注意：给定 n 是一个正整数。 
//
// 示例 1： 
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶 
//
// 示例 2： 
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
// 
// Related Topics 动态规划 
// 👍 1598 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	// (1) --> 1 --> 1
	// (2) --> 1+1 2 --> 2
	// (3) --> (1)+2 (2)+1 --> 3
	// (4) --> (2)+2 (3)+1 --> 5
	// 爬到 n 阶的时候，有两种可能，从 n-2 阶处一次走两阶或者从 n-1 阶处走一阶
	// 故爬到 n 阶只需要关注爬到 n-2 阶处和 n-1 处所需要的方法数
	//dp := make([]int, n+1)
	//dp[1] = 1
	//dp[2] = 2
	//for i := 3; i <= n; i++ {
	//	dp[i] = dp[i-2] + dp[i-1]
	//}
	//return dp[n]
	// 对动态规划进行空间压缩，因为 dp[i] 只与 dp[i-1] 和 dp[i-2] 有关
	if n <= 2 {
		return n
	}
	var cur int
	pre2, pre1 := 1, 2
	for i := 2; i < n; i++ {
		cur = pre2+pre1
		pre2 = pre1
		pre1 = cur
	}
	return cur
}
//leetcode submit region end(Prohibit modification and deletion)
