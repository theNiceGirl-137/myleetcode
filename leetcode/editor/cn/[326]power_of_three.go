//给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。 
//
// 整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 27
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：n = 0
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：n = 9
//输出：true
// 
//
// 示例 4： 
//
// 
//输入：n = 45
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// -231 <= n <= 231 - 1 
// 
//
// 
//
// 进阶： 
//
// 
// 你能不使用循环或者递归来完成本题吗？ 
// 
// Related Topics 递归 数学 
// 👍 170 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	// 有两种方法，一种是从 3 开始以此得出 3 的幂，并与 n 比较大小，一旦大于或者等于停止循环
	i := 1
	for i < n {
		i *= 3
	}
	return i == n
	// 另一种方法是，因为在 int 范围内 3 的最大次方是 3^19=1162261467，如果 n 是 3 的整数次方，那么 1162261467 除以 n 的余数一定是零；反之亦然。
	//return n > 0 && 1162261467 % n == 0
}
//leetcode submit region end(Prohibit modification and deletion)
