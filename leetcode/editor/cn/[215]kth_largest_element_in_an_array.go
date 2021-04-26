//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。 
//
// 示例 1: 
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
// 
//
// 示例 2: 
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4 
//
// 说明: 
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。 
// Related Topics 堆 分治算法 
// 👍 1048 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	// 数组升序排列，第 k 大的数在倒数第 k 位
	target := len(nums)-k
	for l < r {
		// 第 mid 小的数
		mid := quickSelection(nums, l, r)
		if mid == target {
			return nums[mid]
		// 移动区间，因为比第 mid 小的数已经分居左右两边但还未排序
		// 所以移动区间判断下一个数即可
		} else if mid < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}

func quickSelection(nums []int, l, r int) int {
	// l，r 分别是左右边界
	i, j := l+1, r
	for {
		// 将 i 指针右移，找到第一个大于左边界的数
		for i < r && nums[i] <= nums[l] {
			i++
		}
		// 将 j 指针左移，找到第一个小于左边界的数
		for l < j && nums[j] >= nums[l] {
			j--
		}
		// 两指针相遇之后退出，说明大于左边界的数和小于左边界的数各占一边
		if i >= j {
			break
		}
		// 交换两个数的位置
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 此时 j 的位置是第一个小于左边界的数
	nums[l], nums[j] = nums[j], nums[l]
	// 此时 j 前的元素都比 j 小，j 后的元素都比 j 大，j 所在的位置就是第 j 小的数
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
