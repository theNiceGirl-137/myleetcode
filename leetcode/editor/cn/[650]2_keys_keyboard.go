//最初在一个记事本上只有一个字符 'A'。你每次可以对这个记事本进行两种操作： 
//
// 
// Copy All (复制全部) : 你可以复制这个记事本中的所有字符(部分的复制是不允许的)。 
// Paste (粘贴) : 你可以粘贴你上一次复制的字符。 
// 
//
// 给定一个数字 n 。你需要使用最少的操作次数，在记事本中打印出恰好 n 个 'A'。输出能够打印出 n 个 'A' 的最少操作次数。 
//
// 示例 1: 
//
// 
//输入: 3
//输出: 3
//解释:
//最初, 我们只有一个字符 'A'。
//第 1 步, 我们使用 Copy All 操作。
//第 2 步, 我们使用 Paste 操作来获得 'AA'。
//第 3 步, 我们使用 Paste 操作来获得 'AAA'。
// 
//
// 说明: 
//
// 
// n 的取值范围是 [1, 1000] 。 
// 
// Related Topics 动态规划 
// 👍 293 👎 0


package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minSteps(n int) int {
	// 位置 i 表示延展到长度 i 的最少操作次数
	// 对于每个位置 j，如果 j 可以被 i 整除，那么长度 i 就可以由长度 j 操作得到，其操作次数等价于把一个长度为 1 的 A 延展到长度为 i/j（类比于把一个长度为 j 的 A 延展到长度为 i）
	// 因此我们得到的递推公式是 dp[i]=dp[j]+dp[i/j]
	dp := make([]int, n+1)
	h := int(math.Sqrt(float64(n)))
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j <= h; j++ {
			if i % j == 0 {
				//dp[i] = min650(dp[i], dp[j]+dp[i/j])
				dp[i] = dp[j]+dp[i/j]
				break
			}
		}
	}
	return dp[n]
}

func min650(i, j int) int {
	if i < j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
