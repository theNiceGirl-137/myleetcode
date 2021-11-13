//我们定义，在以下情况时，单词的大写用法是正确的： 
//
// 
// 全部字母都是大写，比如 "USA" 。 
// 单词中所有字母都不是大写，比如 "leetcode" 。 
// 如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。 
// 
//
// 给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//输入：word = "USA"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：word = "FlaG"
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= word.length <= 100 
// word 由小写和大写英文字母组成 
// 
// Related Topics 字符串 👍 179 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func detectCapitalUse(word string) bool {
	if allMatch(word[:1], 'a', 'z') {
		return allMatch(word[1:], 'a', 'z')
	}
	if len(word) <= 2 {
		return true
	}
	if allMatch(word[1:2], 'a', 'z') {
		return allMatch(word[2:], 'a', 'z')
	}
	return allMatch(word[2:], 'A', 'Z')
}

func allMatch(s string, l, r byte) bool {
	for i := range s {
		if s[i] < l || s[i] > r {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
