//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。 
//
// 你可以对一个单词进行如下三种操作： 
//
// 
// 插入一个字符 
// 删除一个字符 
// 替换一个字符 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
// 
//
// 示例 2： 
//
// 
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
// 
//
// 
//
// 提示： 
//
// 
// 0 <= word1.length, word2.length <= 500 
// word1 和 word2 由小写英文字母组成 
// 
// Related Topics 字符串 动态规划 
// 👍 1661 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func minDistance72(word1 string, word2 string) int {
	// 类似于题目 1143，使用一个二维数组 dp[i][j]，表示将第一个字符串到位置 i 为止和第二个字符串到 j 为止，最多需要几步编辑
	// 当第 i 位和第 j 位对应的字符相同时，dp[i][j] 等于 dp[i-1][j-1]
	// 当二者对应的字符不相同时，修改的消耗是 dp[i-1][j-1]+1，插入 i 位置/删除 j 位置的消耗是 dp[i][j-1]+1，插入 j 位置/删除 i 位置的消耗是 dp[i-1][j]+1
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				isSame := 0
				if word1[i-1] != word2[j-1] {
					isSame = 1
				}
				dp[i][j] = min72(dp[i-1][j-1]+isSame, min72(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}
	return dp[m][n]
}

func min72(i, j int) int {
	if i < j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
