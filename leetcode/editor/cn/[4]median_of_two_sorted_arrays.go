//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
// 
//
// 示例 2： 
//
// 
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
// 
//
// 示例 3： 
//
// 
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
// 
//
// 示例 4： 
//
// 
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
// 
//
// 示例 5： 
//
// 
//输入：nums1 = [2], nums2 = []
//输出：2.00000
// 
//
// 
//
// 提示： 
//
// 
// nums1.length == m 
// nums2.length == n 
// 0 <= m <= 1000 
// 0 <= n <= 1000 
// 1 <= m + n <= 2000 
// -106 <= nums1[i], nums2[i] <= 106 
// 
//
// 
//
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？ 
// Related Topics 数组 二分查找 分治算法 
// 👍 4020 👎 0


package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1)+len(nums2)
	if l % 2 == 0 {
		return (findKthNumber(nums1, nums2, 0, 0, l/2) + findKthNumber(nums1, nums2, 0, 0, l/2+1)) / 2
	}
	return findKthNumber(nums1, nums2, 0, 0, l/2+1)
}

// l1，l2 是分别指向两个数组的指针
func findKthNumber(nums1, nums2 []int, l1, l2, k int) float64 {
	m1, m2 := 0, 0
	switch {
	// 如果 l1 已经超出了下标，说明只需要从另一个数组中找到剩下的第 k 个数
	case l1 >= len(nums1):
		return float64(nums2[l2+k-1])
	// 如果 l2 已经超出了下标，说明只需要从另一个数组中找到剩下的第 k 个数
	case l2 >= len(nums2):
		return float64(nums1[l1+k-1])
	// 如果已经只剩下一个数要找，立即返回
	case k == 1:
		return float64(min(nums1[l1], nums2[l2]))
	case l1+k/2-1 >= len(nums1):
		m1 = math.MaxInt64
	case l2+k/2-1 >= len(nums2):
		m2 = math.MaxInt64
	default:
		m1 = nums1[l1+k/2-1]
		m2 = nums2[l2+k/2-1]
	}
	if m1 <= m2 {
		return findKthNumber(nums1, nums2, l1+k/2, l2, k-k/2)
	}
	return findKthNumber(nums1, nums2, l1, l2+k/2, k-k/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
