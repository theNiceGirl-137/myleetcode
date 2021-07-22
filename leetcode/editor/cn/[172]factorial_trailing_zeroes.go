//给定一个整数 n，返回 n! 结果尾数中零的数量。 
//
// 示例 1: 
//
// 输入: 3
//输出: 0
//解释: 3! = 6, 尾数中没有零。 
//
// 示例 2: 
//
// 输入: 5
//输出: 1
//解释: 5! = 120, 尾数中有 1 个零. 
//
// 说明: 你算法的时间复杂度应为 O(log n) 。 
// Related Topics 数学 
// 👍 485 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func trailingZeroes(n int) int {
	// 每个尾部的 0 由 2×5=10 而来，因此我们可以把阶乘的每一个元素拆成质数相乘，统计有多少个 2 和 5
	// 明显的，质因子 2 的数量远多于质因子 5 的数量，因此我们可以只统计阶乘结果里有多少个质因子 5
	if n == 0 {
		return 0
	}
	return n/5+trailingZeroes(n/5)
}
//leetcode submit region end(Prohibit modification and deletion)
