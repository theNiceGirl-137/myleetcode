//对于某些固定的 N，如果数组 A 是整数 1, 2, ..., N 组成的排列，使得： 
//
// 对于每个 i < j，都不存在 k 满足 i < k < j 使得 A[k] * 2 = A[i] + A[j]。 
//
// 那么数组 A 是漂亮数组。 
//
// 
//
// 给定 N，返回任意漂亮数组 A（保证存在一个）。 
//
// 
//
// 示例 1： 
//
// 输入：4
//输出：[2,1,4,3]
// 
//
// 示例 2： 
//
// 输入：5
//输出：[3,1,2,5,4] 
//
// 
//
// 提示： 
//
// 
// 1 <= N <= 1000 
// 
//
// 
// Related Topics 数组 数学 分治 
// 👍 131 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func beautifulArray(n int) []int {
	if n == 1 {
		return []int{1}
	}
	// 根据奇偶进行映射
	// 1~N 的值中取出全部奇数，通过 (a+1)/2 映射得到的 1~(N+1)/2 范围的连续值
	odds := beautifulArray((n+1)/2)
	// 返回右侧已排序切片
	evens := beautifulArray(n/2)
	// 根据奇偶进行反射
	for i, v := range odds {
		odds[i] = 2*v-1
	}
	for i, v := range evens {
		evens[i] = 2*v
	}
	// 将左右子问题答案进行拼接
	sli := append(odds, evens...)
	return sli
}
//leetcode submit region end(Prohibit modification and deletion)
