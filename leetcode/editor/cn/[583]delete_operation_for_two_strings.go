//给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。 
//
// 
//
// 示例： 
//
// 输入: "sea", "eat"
//输出: 2
//解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
// 
//
// 
//
// 提示： 
//
// 
// 给定单词的长度不超过500。 
// 给定单词中的字符只含有小写字母。 
// 
// Related Topics 字符串 动态规划 
// 👍 205 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func minDistance583(word1 string, word2 string) int {
	// 定义一个 dp 数组，其中 dp[i] 表示到位置 i 为止的子序列的性质，并不必须以 i 结尾
	// 本题定义一个二维 dp 数组，其中 dp[i][j] 表示到第一个字符串位置 i 为止、到第二个字符串位置 j 为止，最长的公共子序列长度
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return m+n-2*dp[m][n]
}
//leetcode submit region end(Prohibit modification and deletion)
