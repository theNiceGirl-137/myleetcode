//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。 
//
// 示例: 
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
// Related Topics 回溯算法 
// 👍 579 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	var ans [][]int
	comb := make([]int, k)
	count := 0
	backTracking77(&ans, comb, count, 1, n, k)
	return ans
}

func backTracking77(ans *[][]int, comb []int, count, pos, n, k int) {
	if count == k {
		temp := make([]int, len(comb))
		copy(temp, comb)
		*ans = append(*ans, temp)
		return
	}
	for i := pos; i <= n; i++ {
		comb[count] = i
		count++
		backTracking77(ans, comb, count, i+1, n, k)
		count--
	}
}
//leetcode submit region end(Prohibit modification and deletion)
