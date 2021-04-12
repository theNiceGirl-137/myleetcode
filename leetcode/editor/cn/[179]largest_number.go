//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。 
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [10,2]
//输出："210" 
//
// 示例 2： 
//
// 
//输入：nums = [3,30,34,5,9]
//输出："9534330"
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出："1"
// 
//
// 示例 4： 
//
// 
//输入：nums = [10]
//输出："10"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 109 
// 
// Related Topics 排序 
// 👍 575 👎 0


package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		a, _ := strconv.Atoi(strs[i]+strs[j])
		b, _ := strconv.Atoi(strs[j]+strs[i])
		return a > b
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
//leetcode submit region end(Prohibit modification and deletion)
