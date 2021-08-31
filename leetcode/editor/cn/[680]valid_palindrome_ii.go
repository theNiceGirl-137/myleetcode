//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。 
//
// 示例 1: 
//
// 
//输入: "aba"
//输出: True
// 
//
// 示例 2: 
//
// 
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
// 
//
// 注意: 
//
// 
// 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。 
// 
// Related Topics 字符串 
// 👍 342 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			// 右边的字符向右移动一位，左边的字符向左移动一位，相当于移除左边或者右边的一个字符
			// 区间左闭右开
			return isPalindrome680(s[l:r]) || isPalindrome680(s[l+1:r+1])
		}
		l++
		r--
	}
	return true
}

func isPalindrome680(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
