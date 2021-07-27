//给你一个 正 整数 num ，输出它的补数。补数是对该数的二进制表示取反。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：num = 5
//输出：2
//解释：5 的二进制表示为 101（没有前导零位），其补数为 010。所以你需要输出 2 。
// 
//
// 示例 2： 
//
// 
//输入：num = 1
//输出：0
//解释：1 的二进制表示为 1（没有前导零位），其补数为 0。所以你需要输出 0 。
// 
//
// 
//
// 提示： 
//
// 
// 给定的整数 num 保证在 32 位带符号整数的范围内。 
// num >= 1 
// 你可以假定二进制数不包含前导零位。 
// 本题与 1009 https://leetcode-cn.com/problems/complement-of-base-10-integer/ 相同 
// 
// Related Topics 位运算 
// 👍 216 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func findComplement(num int) int {
	var ans int
	n, cnt := num, 0
	// 统计位数
	for n > 0 {
		cnt++
		n >>= 1
	}
	// a ^ 1 = ~a，所以将ans置为全 1
	for i := 0; i < cnt; i++ {
		ans <<= 1
		ans += 1
	}
	return ans ^ num
}
//leetcode submit region end(Prohibit modification and deletion)
