//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的
//房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。 
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [2,3,2]
//输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3,1]
//输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 1000 
// 
// Related Topics 数组 动态规划 
// 👍 722 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func rob213(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 环形数组意味着头尾只能选择一间
	return max213(maxProfile(nums[1:]), maxProfile(nums[:len(nums)-1]))
}

func maxProfile(nums []int) int {
	//dp := make([]int, len(nums)+1)
	//dp[1] = nums[0]
	//for i := 2; i <= len(nums); i++ {
	//	dp[i] = max213(dp[i-1], dp[i-2]+nums[i-1])
	//}
	//return dp[len(nums)]
	var cur int
	pre2, pre1 := 0, 0
	for i := 0; i < len(nums); i++ {
		cur = max213(pre1, pre2+nums[i])
		pre2 = pre1
		pre1 = cur
	}
	return cur
}

func max213(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
