//给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：3
//解释：
//只需要3次操作（注意每次操作会增加两个元素的值）：
//[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,1,1]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// n == nums.length 
// 1 <= nums.length <= 10⁵ 
// -10⁹ <= nums[i] <= 10⁹ 
// 答案保证符合 32-bit 整数 
// 
// Related Topics 数组 数学 👍 304 👎 0


package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func minMoves(nums []int) int {
	// 若只关注数字之间的差，n-1 个数都加一相当于某一个数减一
	var ans int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		ans += nums[i]-nums[0]
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
