//有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。 
//
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i -
// 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。 
//
// 求所能获得硬币的最大数量。 
//
// 
//示例 1：
//
// 
//输入：nums = [3,1,5,8]
//输出：167
//解释：
//nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
//coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167 
//
// 示例 2： 
//
// 
//输入：nums = [1,5]
//输出：10
// 
//
// 
//
// 提示： 
//
// 
// n == nums.length 
// 1 <= n <= 500 
// 0 <= nums[i] <= 100 
// 
// Related Topics 数组 动态规划 
// 👍 761 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	m := len(nums)
	// 在数组前后补充 1
	temp := make([]int, m+2)
	temp[0], temp[m+1] = 1, 1
	for i := 0; i < m; i++ {
		temp[i+1] = nums[i]
	}
	// dp[i][j] 表示 i～j 区间（不包括 i 和 j）戳破能获得的最大值
	dp := make([][]int, m+2)
	for i := range dp {
		dp[i] = make([]int, m+2)
	}
	// i 表示右边界，因为至少三个气球才会有区间，所以从 3 开始
	for i := 3; i <= m+2; i++ {
		// j 表示左边界
		for j := 0; j < m+3-i; j++ {
			var res int
			// k 表示 k 左右的气球都被戳破的情况下所能获得的最大值
			for k := j+1; k < j+i-1; k++ {
				left := dp[j][k]
				right := dp[k][j+i-1]
				res = max312(res, left+temp[j]*temp[k]*temp[j+i-1]+right)
			}
			dp[j][j+i-1] = res
		}
	}
	return dp[0][m+1]
}

func max312(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
