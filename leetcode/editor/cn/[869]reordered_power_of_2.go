//给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。 
//
// 如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入：1
//输出：true
// 
//
// 示例 2： 
//
// 输入：10
//输出：false
// 
//
// 示例 3： 
//
// 输入：16
//输出：true
// 
//
// 示例 4： 
//
// 输入：24
//输出：false
// 
//
// 示例 5： 
//
// 输入：46
//输出：true
// 
//
// 
//
// 提示： 
//
// 
// 1 <= N <= 10^9 
// 
// Related Topics 数学 计数 枚举 排序 👍 96 👎 0


package leetcode

import (
	"sort"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func reorderedPowerOf2(n int) bool {
	// 枚举数组排列，判断每个不含前导 0 的排列是否是 2 的幂
	nums := []byte(strconv.Itoa(n))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return backTracking869(0, 0, len(nums), nums, make([]bool, len(nums)))
}

func backTracking869(idx, num, m int, nums []byte, visited []bool) bool {
	if idx == m {
		return isPowerOf2(num)
	}
	for i, ch := range nums {
		// 排除前导 0
		// 如果与前一个数相同，说明前一个数已经选择完了可能的组合，直接跳过
		if num == 0 && ch == '0' || visited[i] || i > 0 && !visited[i-1] && ch == nums[i-1] {
			continue
		}
		visited[i] = true
		if backTracking869(idx+1, num*10+int(ch-'0'), m, nums, visited) {
			return true
		}
		visited[i] = false
	}
	return false
}

func isPowerOf2(n int) bool {
	return n&(n-1) == 0
}
//leetcode submit region end(Prohibit modification and deletion)
