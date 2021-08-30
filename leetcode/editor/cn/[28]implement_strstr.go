//实现 strStr() 函数。 
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如
//果不存在，则返回 -1 。 
//
// 
//
// 说明： 
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。 
//
// 
//
// 示例 1： 
//
// 
//输入：haystack = "hello", needle = "ll"
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1
// 
//
// 示例 3： 
//
// 
//输入：haystack = "", needle = ""
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= haystack.length, needle.length <= 5 * 10⁴ 
// haystack 和 needle 仅由小写英文字符组成 
// 
// Related Topics 双指针 字符串 字符串匹配 👍 1012 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	// KMP 算法
	k, n, p := -1, len(haystack), len(needle)
	if p == 0 {
		return 0
	}
	next := make([]int, p)
	for i := 0; i < len(next); i++ {
		// -1 表示不存在最大的前缀和后缀
		next[i] = -1
	}
	calNext(needle, next)
	for i := 0; i < n; i++ {
		for k > -1 && needle[k+1] != haystack[i] {
			// 有部分匹配，往前回溯
			k = next[k]
		}
		if needle[k+1] == haystack[i] {
			k++
		}
		if k == p-1 {
			// 说明 k 移动到 needle 最末端，返回相应的位置
			return i-p+1
		}
	}
	return -1
}

// calNext 计算 next 数组
func calNext(needle string, next []int) {
	// 本质是动态规划
	for j, p := 1, -1; j < len(needle); j++ {
		for p > -1 && needle[p+1] != needle[j] {
			// 如果下一位不同，往前回溯
			p = next[p]
		}
		if needle[p+1] == needle[j] {
			// 如果下一位相同，更新相同的最大前缀和最大后缀长
			p++
		}
		next[j] = p
	}
}
//leetcode submit region end(Prohibit modification and deletion)
