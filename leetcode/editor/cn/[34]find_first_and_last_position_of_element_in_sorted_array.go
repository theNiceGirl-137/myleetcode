//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。 
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。 
//
// 进阶： 
//
// 
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？ 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4] 
//
// 示例 2： 
//
// 
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1] 
//
// 示例 3： 
//
// 
//输入：nums = [], target = 0
//输出：[-1,-1] 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 105 
// -109 <= nums[i] <= 109 
// nums 是一个非递减数组 
// -109 <= target <= 109 
// 
// Related Topics 数组 二分查找 
// 👍 966 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		// 注意退出条件
		mid := (r+l)/2
		if nums[mid] == target {
			i := mid
			j := mid
			for i >= 0 && nums[i] == target {
				i--
			}
			for j <= len(nums)-1 && nums[j] == target {
				j++
			}
			return []int{i + 1, j - 1}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}
//leetcode submit region end(Prohibit modification and deletion)
