//给出两个整数 n 和 k，找出所有包含从 1 到 n 的数字，且恰好拥有 k 个逆序对的不同的数组的个数。 
//
// 逆序对的定义如下：对于数组的第i个和第 j个元素，如果满i < j且 a[i] > a[j]，则其为一个逆序对；否则不是。 
//
// 由于答案可能很大，只需要返回 答案 mod 10⁹ + 7 的值。 
//
// 示例 1: 
//
// 
//输入: n = 3, k = 0
//输出: 1
//解释: 
//只有数组 [1,2,3] 包含了从1到3的整数并且正好拥有 0 个逆序对。
// 
//
// 示例 2: 
//
// 
//输入: n = 3, k = 1
//输出: 2
//解释: 
//数组 [1,3,2] 和 [2,1,3] 都有 1 个逆序对。
// 
//
// 说明: 
//
// 
// n 的范围是 [1, 1000] 并且 k 的范围是 [0, 1000]。 
// 
// Related Topics 动态规划 👍 163 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func kInversePairs(n int, k int) int {
	dp := make([]int, k+1)
	// 维护前缀和
	sum := make([]int, k+2)
	dp[0] = 1
	sum[1]=1
	for i := 1; i <= n; i++ {
		// 对于 n 来说有 n-1 个逆序对，对于 n-1 来说有 n-2 个逆序对，...，对于 1 来说没有逆序对
		// 最多有 (n-1) + (n-2) + ... + 0 = n(n-1)/2，对于 k 大于 n(n-1)/2 的输入直接返回 0
		for j := min(k, i*(i-1)/2); j > 0; j-- {
			dp[j] = sum[j+1]-sum[max(0, j-i+1)]
			dp[j] %= 1e9+7
		}
		for j := 1; j <= k; j++ {
			sum[j+1] = sum[j]+dp[j]
		}
	}
	return dp[k]
}
//leetcode submit region end(Prohibit modification and deletion)
