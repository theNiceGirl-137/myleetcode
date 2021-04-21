//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。 
//
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
// 
//
// 示例 2： 
//
// 
//输入：s = "a", t = "a"
//输出："a"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length, t.length <= 105 
// s 和 t 由英文字母组成 
// 
//
// 
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？ Related Topics 哈希表 双指针 字符串 Sliding Window 
// 👍 1120 👎 0



package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	need := make(map[byte]int, 0)
	// 统计 t 中的字符情况
	for i := range t {
		need[t[i]]++
	}
	// l 和 r 作为窗口的左右指针，count 作为已经加入的元素的数量，当 count 等于 t 的长度时表示
	// 当前窗口中包含了全部需要的元素
	l, r, count, minL, minSize := 0, 0, 0, 0, len(s)+1
	for r < len(s) {
		if v, ok := need[s[r]]; ok && v-1 >= 0 {
			count++
			need[s[r]]--
			// 若当前窗口已经包含 t 中全部的字符
			// 移动 l，在不影响结果的情况下获得最短的字符串
			for count == len(t) {
				if r-l+1 < minSize {
					minL = l
					minSize = r-l+1
				}
				if v, ok = need[s[l]]; ok && v+1 > 0 {
					need[s[l]]++
					count--
				}
				l++
			}
		}
		r++
	}
	if minSize > len(s) {
		return ""
	}
	return s[minL:minL+minSize]
}
//leetcode submit region end(Prohibit modification and deletion)
