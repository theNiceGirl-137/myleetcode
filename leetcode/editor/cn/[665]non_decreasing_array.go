//给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。 
//
// 我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [4,2,3]
//输出: true
//解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
// 
//
// 示例 2: 
//
// 
//输入: nums = [4,2,1]
//输出: false
//解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 10 ^ 4 
// - 10 ^ 5 <= nums[i] <= 10 ^ 5 
// 
// Related Topics 数组 
// 👍 567 👎 0


package main
//leetcode submit region begin(Prohibit modification and deletion)
func checkPossibility(nums []int) bool {
	// 如果将 nums[i] 缩小，可能会导致其无法融入前面已经遍历过的非递减子数列
	// 如果将 nums[i+1] 放大，可能会导致其后续的继续出现递减
	if len(nums) == 1 {
		return true
	}
	flag := (nums[0] <= nums[1])
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if flag {
				// 需要尽可能不放大 nums[i+1]，这样会让后续非递减更困难
				// 如果缩小 nums[i]，但不破坏前面的子序列的非递减性
				if nums[i+1] >= nums[i-1] {
					nums[i] = nums[i+1]
				} else {
					nums[i+1] = nums[i]
				}
				flag = false
			} else {
				return false
			}
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
