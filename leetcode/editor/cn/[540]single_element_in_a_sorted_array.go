//给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。 
//
// 示例 1: 
//
// 
//输入: [1,1,2,3,3,4,4,8,8]
//输出: 2
// 
//
// 示例 2: 
//
// 
//输入: [3,3,7,7,10,11,11]
//输出: 10
// 
//
// 注意: 您的方案应该在 O(log n)时间复杂度和 O(1)空间复杂度中运行。 
// Related Topics 二分查找 
// 👍 229 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l+r)/2
		// 出现独立数字之前，相等的两个数为先偶后奇，出现之后为先奇后偶
		// 下标为奇数，和左右两边的数对比
		if m % 2 == 1 {
			// 说明现在在出现独立数字之前
			if m > 0 && nums[m] == nums[m-1] {
				l = m + 1
			}
			// 说明现在出现独立数字之后
			if m > 0 && m + 1 < len(nums) && nums[m] == nums[m+1] {
				r = m - 1
			}
		} else {
			// 独立数字一定在偶数位
			// 下标为偶数，和左右两边的数对比
			// 说明现在在出现独立数字之前，直接移动到下一个不重复数字的位置上
			if m + 1 < len(nums) && nums[m] == nums[m+1] {
				l = m + 2
			// 说明现在在出现独立数字之后，直接移动到下一个不重复数字的位置上
			} else if m > 0 && nums[m] == nums[m-1] {
				r = m - 2
			} else {
				// 和左右两边都不相等，找到独立的数
				return nums[m]
			}
		}
	}
	return nums[l]
}
//leetcode submit region end(Prohibit modification and deletion)
