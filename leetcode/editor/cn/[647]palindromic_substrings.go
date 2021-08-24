//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。 
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。 
//
// 
//
// 示例 1： 
//
// 输入："abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
// 
//
// 示例 2： 
//
// 输入："aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa" 
//
// 
//
// 提示： 
//
// 
// 输入的字符串长度不会超过 1000 。 
// 
// Related Topics 字符串 动态规划 
// 👍 658 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		// 数组长度可能为奇数也可能为偶数
		count += extendSubstring(s, i, i)
		count += extendSubstring(s, i, i+1)
	}
	return count
}

func extendSubstring(s string, l, r int) int {
	var count int
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向左右两边延伸
		l--
		r++
		count++
	}
	return count
}
//leetcode submit region end(Prohibit modification and deletion)
