//统计所有小于非负整数 n 的质数的数量。 
//
// 
//
// 示例 1： 
//
// 输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
// 
//
// 示例 2： 
//
// 输入：n = 0
//输出：0
// 
//
// 示例 3： 
//
// 输入：n = 1
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= n <= 5 * 106 
// 
// Related Topics 数组 数学 枚举 数论 
// 👍 718 👎 0


package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	// 埃拉托斯特尼筛法
	// 从 1 到 n 遍历，假设当前遍历到 m，则把所有小于 n 的、且是 m 的倍数的整数标为和数
	// 遍历完成后，没有被标为和数的数字即为质数
	if n <= 2 {
		return 0
	}
	//num := n-2 // 去掉 1 和 2
	num := n/2 // 偶数一定不是质数，1 跟 2 相反
	prime := make([]bool, n+1)
	for i := range prime {
		prime[i] = true
	}
	i := 3
	sqrtn := int(math.Sqrt(float64(n))) // 最小质因子一定小于开方数
	for i <= sqrtn {
		for j := i*i; j < n; j += 2*i {
			if prime[j] {
				prime[j] = false
				num--
			}
		}
		// 避免偶数和重复遍历
		for i += 2; i <= sqrtn && !prime[i]; {
			i += 2
		}
	}
	//for i := 2; i <= n; i++ {
	//	if prime[i] {
	//		for j := 2*i; j < n; j += i {
	//			if prime[j] {
	//				prime[j] = false
	//				num--
	//			}
	//		}
	//	}
	//}
	return num
}
//leetcode submit region end(Prohibit modification and deletion)
