//给你一个字符串 s，找到 s 中最长的回文子串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
// 
//
// 示例 2： 
//
// 
//输入：s = "cbbd"
//输出："bb"
// 
//
// 示例 3： 
//
// 
//输入：s = "a"
//输出："a"
// 
//
// 示例 4： 
//
// 
//输入：s = "ac"
//输出："a"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 仅由数字和英文字母（大写和/或小写）组成 
// 
// Related Topics 字符串 动态规划 👍 4034 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	length := len(s)
	maxStart := 0   // 最长回文串的起点
	maxEnd := 0 // 最长回文串的终点
	maxLen := 1 // 最长回文串的长度

	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	// 这个是求左右边界位置，固定的模板，r 从 1 开始，l 从 r 往前开始
	for r := 1; r < length; r++ {
		for l := 0; l < r; l++ {
			// 2 个字符串相等，去看上一个 dp 或是否是相邻的情况
			if s[l] == s[r] && (r - l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r - l + 1 > maxLen {
					maxLen = r - l + 1
					maxStart = l
					maxEnd = r
				}
			}
		}
	}
	return s[maxStart:maxEnd+1]
}
//leetcode submit region end(Prohibit modification and deletion)
