//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。 
//
// 说明：每次只能向下或者向右移动一步。 
//
// 
//
// 示例 1： 
//
// 
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
// 
//
// 示例 2： 
//
// 
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 200 
// 0 <= grid[i][j] <= 100 
// 
// Related Topics 数组 动态规划 
// 👍 884 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	// 定义一个同样是二维的 dp 数组，其中 dp[i][j] 表示从左上角开始到 (i, j) 位置的最优路径数字和
	// 因为每次只能向下或者向右移动，所以状态转移方程为 dp[i][j]=min(dp[i-1][j], dp[i][j-1])+grid[i][j]
	//dp := make([][]int, len(grid))
	//for i := 0; i < len(grid); i++ {
	//	dp[i] = make([]int, len(grid[0]))
	//}
	//for i := 0; i < len(grid); i++ {
	//	for j := 0; j < len(grid[0]); j++ {
	//		if i == 0 && j == 0 {
	//			dp[i][j] = grid[i][j]
	//		} else if i == 0 {
	//			dp[i][j] = dp[i][j-1]+grid[i][j]
	//		} else if j == 0 {
	//			dp[i][j] = dp[i-1][j]+grid[i][j]
	//		} else {
	//			dp[i][j] = min(dp[i-1][j], dp[i][j-1])+grid[i][j]
	//		}
	//	}
	//}
	//return dp[len(grid)-1][len(grid[0])-1]
	// 因为 dp 矩阵的每一个值只和左边和上面的值相关，使用空间压缩将 dp 数组压缩为一维
	// 对于第 i 行，在遍历到第 j 列的时候，因为第 j-1 列已经更新过了，所以 dp[j-1] 代表 dp[i][j-1]的值
	// 而 dp[j] 待更新，当前存储的值是第 i-1 行的时候计算的，所以代表 dp[i-1][j] 的值
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1]+grid[i][j]
			} else if j == 0 {
				dp[j] = dp[j]+grid[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1])+grid[i][j]
			}
		}
	}
	return dp[len(grid[0])-1]
}
//leetcode submit region end(Prohibit modification and deletion)
