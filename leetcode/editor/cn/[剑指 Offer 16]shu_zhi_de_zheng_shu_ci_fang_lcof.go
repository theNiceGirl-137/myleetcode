//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xⁿ）。不得使用库函数，同时不需要考虑大数问题。 
//
// 
//
// 示例 1： 
//
// 
//输入：x = 2.00000, n = 10
//输出：1024.00000
// 
//
// 示例 2： 
//
// 
//输入：x = 2.10000, n = 3
//输出：9.26100 
//
// 示例 3： 
//
// 
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2⁻² = 1/2² = 1/4 = 0.25 
//
// 
//
// 提示： 
//
// 
// -100.0 < x < 100.0 
// -2³¹ <= n <= 2³¹-1 
// -10⁴ <= xⁿ <= 10⁴ 
// 
//
// 
//
// 注意：本题与主站 50 题相同：https://leetcode-cn.com/problems/powx-n/ 
// Related Topics 递归 数学 👍 209 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	// 快速幂可以用位运算来实现
	// n&1 取n的二进制数最低位，n&1=1，n为奇数，等于 0 为偶数，相当于 n%2=0
	// n>>1 右移1位，去掉 n 的二进制数最低位，相当于 n/2
	// 当 n 为偶数时
	// x^n=x^(n/2)*x^(n/2)
	// n>>1，n 右移 1 位后，x 自己乘自己，因为二进制每位的差距是平方关系
	// 当 n 为奇数时
	// 需要再乘以多出来的一次，即 x^n=x*x^(n-1)
	// n-1，x 不更新，将 x 累乘到结果
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1/x
		n = -n
	}
	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp*temp
	}
	return x*temp*temp
}
//leetcode submit region end(Prohibit modification and deletion)
