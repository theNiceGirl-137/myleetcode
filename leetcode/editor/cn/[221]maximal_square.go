//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。 
//
// 
//
// 示例 1： 
//
// 
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
// 
//
// 示例 2： 
//
// 
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
// 
//
// 示例 3： 
//
// 
//输入：matrix = [["0"]]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 300 
// matrix[i][j] 为 '0' 或 '1' 
// 
// Related Topics 动态规划 
// 👍 776 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	// 对于在矩阵内搜索正方形或者长方形的题目，一种常见的做法是定义一个二维 dp 数组
	// 其中 dp[i][j] 表示满足题目条件的、以 (i, j) 为右下角的正方形或者长方形的属性
	// 对于本题则表示以 (i, j) 为右下角的的全由 1 构成的最大正方形的面积
	// 如果当前位置是 0，那么 dp[i][j] 即为 0
	// 如果当前位置是 1，那么假设 dp[i][j]=k，其充分条件为 dp[i-1][j-1]、dp[i][j-1] 和 dp[i-1][j] 的值必须是
	// 都不小于 k-1，否则 (i, j) 位置不可以构成一个边长为 k 的正方形
	// 同理，如果这三个值的最小值为 k-1，则 (i, j) 位置一定且最大可以构成一个边长为 k 的正方形
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return 0
	}
	var maxSide int
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))+1
			}
			maxSide = max(maxSide, dp[i][j])
		}
	}
	return maxSide*maxSide
}
//leetcode submit region end(Prohibit modification and deletion)
