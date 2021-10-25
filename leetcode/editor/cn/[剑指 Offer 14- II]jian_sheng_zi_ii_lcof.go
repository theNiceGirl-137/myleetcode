//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1]
// 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘
//积是18。 
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。 
//
// 
//
// 示例 1： 
//
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1 
//
// 示例 2: 
//
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36 
//
// 
//
// 提示： 
//
// 
// 2 <= n <= 1000 
// 
//
// 注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/ 
// Related Topics 数学 动态规划 👍 137 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func cuttingRopeOffer14_2(n int) int {
	// ① 当所有绳段长度相等时，乘积最大
	// ② 最优的绳段长度为 3
	// 证明过程：https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/
	if n <= 3 {
		return n-1
	}
	// n=3a+b
	res := 1
	for n > 4 {
		res = res*3%1000000007
		n -= 3
	}
	// 最后结果只会剩下 2,3,4，所以直接乘 n 再取余
	// 因为 2、3 不需要再剪了（剪了反而变小）
	// 4 剪成 2x2 是最大的，2x2 恰巧等于 4
	return res*n%1000000007
}
//leetcode submit region end(Prohibit modification and deletion)
