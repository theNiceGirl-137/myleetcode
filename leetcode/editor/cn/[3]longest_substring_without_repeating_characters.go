//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 
//
// 示例 2: 
//
// 
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 
//
// 示例 3: 
//
// 
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
// 示例 4: 
//
// 
//输入: s = ""
//输出: 0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 5 * 10⁴ 
// s 由英文字母、数字、符号和空格组成 
// 
// Related Topics 哈希表 字符串 滑动窗口 👍 6023 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var left, right, ans int
	m := make(map[byte]int)
	for right < len(s) {
		//if _, ok := m[s[right]]; !ok {
		//	m[s[right]] = right
		//} else {
		//	if m[s[right]]+1 >= left {
		//		left = m[s[right]]+1
		//	}
		//	m[s[right]] = right
		//}
		if _, ok := m[s[right]]; ok && m[s[right]]+1 >= left {
			left = m[s[right]]+1
		}
		m[s[right]] = right
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
