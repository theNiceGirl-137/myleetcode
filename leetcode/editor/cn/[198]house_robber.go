//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。 
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。 
//
// 
//
// 示例 1： 
//
// 
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。 
//
// 示例 2： 
//
// 
//输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 400 
// 
// Related Topics 动态规划 
// 👍 1484 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func rob198(nums []int) int {
	// 定义一个数组 dp，dp[i] 表示抢劫到第 i 个房子时可以抢劫的最大数量
	// 对于 dp[i]，一是选择不抢这个房子，此时累计的金额为 dp[i-1]
	// 二是选择抢这个房子，那么累计的最大金额只能是 dp[i-2]，因为如果抢第 i-1 个房子会出发警报
	// 故状态转移方程为 dp[i]=max(dp[i-1], nums[i-1]+dp[i-2])
	//if len(nums) < 1 {
	//	return 0
	//}
	//dp := make([]int, len(nums)+1)
	//dp[1] = nums[0]
	//for i := 2; i <= len(nums); i++ {
	//	dp[i] = max198(dp[i-1], dp[i-2]+nums[i-1])
	//}
	//return dp[len(nums)]
	// 对空间进行压缩
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var cur int
	pre2, pre1 := 0, 0
	for i := 0; i < len(nums); i++ {
		cur = max198(pre1, pre2+nums[i])
		pre2 = pre1
		pre1 = cur
	}
	return cur
}

func max198(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
