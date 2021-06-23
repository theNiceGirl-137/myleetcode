//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。 
//
// 
// '.' 匹配任意单个字符 
// '*' 匹配零个或多个前面的那一个元素 
// 
//
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。 
// 
//
// 示例 1： 
//
// 
//输入：s = "aa" p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。
// 
//
// 示例 2: 
//
// 
//输入：s = "aa" p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
// 
//
// 示例 3： 
//
// 
//输入：s = "ab" p = ".*"
//输出：true
//解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
// 
//
// 示例 4： 
//
// 
//输入：s = "aab" p = "c*a*b"
//输出：true
//解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
// 
//
// 示例 5： 
//
// 
//输入：s = "mississippi" p = "mis*is*p*."
//输出：false 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 20 
// 0 <= p.length <= 30 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。 
// 保证每次出现字符 * 时，前面都匹配到有效的字符 
// 
// Related Topics 字符串 动态规划 回溯算法 
// 👍 2192 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	// dp[i][j] 表示以 i 截止的字符串是否可以被以 j 截止的字符串正则表达式匹配
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	// '*' 匹配零个或多个前面的那一个元素
	// '.' 匹配任意单个字符
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			// 处理 '.'
			if p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			// 处理字符
			} else if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] && p[j-1] == s[i-1]
			// 处理 '*'
			} else if p[j-2] != s[i-1] && p[j-2] != '.' {
				dp[i][j] = dp[i][j-2]
			} else {
				// dp[i][j] = dp[i-1][j] 多个字符匹配的情况
				// dp[i][j] = dp[i][j-1] 单个字符匹配的情况
				// dp[i][j] = dp[i][j-2] 没有匹配的情况
				dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
			}
		}
	}
	return dp[m][n]
}
//leetcode submit region end(Prohibit modification and deletion)
