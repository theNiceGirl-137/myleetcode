//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。 
//
// 
// 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。 
//
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
//输出：4
//解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
//其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 
//n 的值 3 。
// 
//
// 示例 2： 
//
// 
//输入：strs = ["10", "0", "1"], m = 1, n = 1
//输出：2
//解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= strs.length <= 600 
// 1 <= strs[i].length <= 100 
// strs[i] 仅由 '0' 和 '1' 组成 
// 1 <= m, n <= 100 
// 
// Related Topics 动态规划 
// 👍 512 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
// 统计当前字符串内 0 和 1 的数量
type pair struct {
	count0, count1 int
}

func findMaxForm(strs []string, m int, n int) int {
	// 多维费用 0-1 背包问题，有两个背包大小，0 的数量和 1 的数量
	// 数组的 i 和 j 分别表示 0 和 1 的个数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		p := count(str)
		for i := m; i >= p.count0; i-- {
			for j := n; j >= p.count1; j-- {
				// 含有当前 0 1 个数的字符串子集，最多的子串个数为
				// 一、不加入当前字符串
				// 二、加入当前字符串
				dp[i][j] = max474(dp[i][j], dp[i-p.count0][j-p.count1]+1)
			}
		}
	}
	return dp[m][n]
}

// 计算字符串中 0 1 个数的函数
func count(s string) pair {
	count0, count1 := len(s), 0
	for _, char := range s {
		if char == '1' {
			count1++
			count0--
		}
	}
	return pair{count0: count0, count1: count1}
}

func max474(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
