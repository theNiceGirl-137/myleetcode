//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。 
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序
//列。 
// 
//
// 示例 1： 
//
// 
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1,0,3,2,3]
//输出：4
// 
//
// 示例 3： 
//
// 
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2500 
// -104 <= nums[i] <= 104 
// 
//
// 
//
// 进阶： 
//
// 
// 你可以设计时间复杂度为 O(n2) 的解决方案吗？ 
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗? 
// 
// Related Topics 二分查找 动态规划 
// 👍 1646 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	//maxLen := 1
	// dp[i] 表示以 i 结尾的最长子序列长度
	//dp := make([]int, len(nums))
	//for i := range dp {
	//	dp[i] = 1
	//}
	//for i := 0; i < len(nums); i++ {
	//	for j := 0; j < i; j++ {
	//		if nums[i] > nums[j] {
	//			dp[i] = max(dp[i], dp[j]+1)
	//		}
	//	}
	//	maxLen = max(dp[i], maxLen)
	//}
	//return maxLen
	// dp[i] 存储长度为 i+1 的最长递增子序列的最后一个数字
	dp := make([]int, 0)
	dp = append(dp, nums[0])
	// 遍历每一个位置 i，如果其对应的数字大于 dp 中的所有值，直接放在 dp 数组尾部
	// 如果发现这个数字在 dp 数组中比 a 大，比 b 小，则将 b 更新为此数字，使得此后递增序列的可能性变大
	for i := 1; i < len(nums); i++ {
		if dp[len(dp)-1] < nums[i] {
			dp = append(dp, nums[i])
		} else {
			index := lowerBound(0, len(dp)-1, dp, nums[i])
			if index != -1 {
				dp[index] = nums[i]
			}
		}
	}
	return len(dp)
}

func lowerBound(start, end int, nums []int, target int) int {
	for ; start <= end; start++ {
		if nums[start] >= target {
			return start
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
