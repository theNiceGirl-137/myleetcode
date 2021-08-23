//给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。 
//
// 返回 A 的任意排列，使其相对于 B 的优势最大化。 
//
// 
//
// 示例 1： 
//
// 输入：A = [2,7,11,15], B = [1,10,4,11]
//输出：[2,11,7,15]
// 
//
// 示例 2： 
//
// 输入：A = [12,24,8,32], B = [13,25,32,11]
//输出：[24,32,8,12]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= A.length = B.length <= 10000 
// 0 <= A[i] <= 10^9 
// 0 <= B[i] <= 10^9 
// 
// Related Topics 贪心 数组 排序 
// 👍 143 👎 0


package leetcode

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func advantageCount(nums1 []int, nums2 []int) []int {
	// 田忌赛马
	pos := make([][2]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		pos[i] = [2]int{nums2[i], i}
	}
	sort.Ints(nums1)
	sort.Slice(pos, func(i, j int) bool {
		return pos[i][0] < pos[j][0]
	})
	var temp int
	for i := len(pos)-1; i >= 0; i-- {
		if nums1[i+temp] > pos[i][0] {
			nums2[pos[i][1]] = nums1[i+temp]
		} else {
			nums2[pos[i][1]] = nums1[temp]
			temp++
		}
	}
	return nums2
}
//leetcode submit region end(Prohibit modification and deletion)
