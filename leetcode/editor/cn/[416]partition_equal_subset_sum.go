//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 200 
// 1 <= nums[i] <= 100 
// 
// Related Topics 动态规划 
// 👍 832 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	// 0-1 背包问题
	var sum, target, n int
	for _, v := range nums {
		sum += v
	}
	if sum % 2 == 1 {
		return false
	}
	target, n = sum/2, len(nums)
	//dp := make([][]bool, n+1)
	//for i := 0; i <= n; i++ {
	//	dp[i] = make([]bool, target+1)
	//	dp[i][0] = true
	//}
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= target; j++ {
	//		if j-nums[i-1] < 0 {
	//			dp[i][j] = dp[i-1][j]
	//		} else {
	//			dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
	//		}
	//	}
	//}
	//return dp[n][target]
	// 进行空间压缩，注意对数字的遍历需要逆向
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := target; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
	}
	return dp[target]
}
//leetcode submit region end(Prohibit modification and deletion)
