//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第
//一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。 
//
// 示例 1: 
//
// 
//输入: [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数； 
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
// 
//
// 注意: 输入数组的长度不会超过 10000。 
// Related Topics 栈 数组 单调栈 
// 👍 470 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	// 扩容数组，模拟循环数组，因为最多只需要重新遍历一遍数组
	nums = append(nums, nums...)
	n := len(nums)
	ans := make([]int, n)
	indices := make([]int, 0)
	for i := 0; i < n; i++ {
		// 单调栈，栈内元素始终保持递减
		for len(indices) != 0 {
			preIndex := indices[len(indices)-1]
			if nums[i] <= nums[preIndex] {
				break
			}
			indices = indices[:len(indices)-1]
			ans[preIndex] = nums[i]
		}
		// 存放位置
		indices = append(indices, i)
	}
	for i := 0; i < len(indices); i++ {
		ans[indices[i]] = -1
	}
	return ans[:n/2]
}
//leetcode submit region end(Prohibit modification and deletion)
