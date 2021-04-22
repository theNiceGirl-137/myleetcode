//给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，返回长度最长且字典顺序最小的字符
//串。如果答案不存在，则返回空字符串。 
//
// 示例 1: 
//
// 
//输入:
//s = "abpcplea", d = ["ale","apple","monkey","plea"]
//
//输出: 
//"apple"
// 
//
// 示例 2: 
//
// 
//输入:
//s = "abpcplea", d = ["a","b","c"]
//
//输出: 
//"a"
// 
//
// 说明: 
//
// 
// 所有输入的字符串只包含小写字母。 
// 字典的大小不会超过 1000。 
// 所有输入的字符串长度不会超过 1000。 
// 
// Related Topics 排序 双指针 
// 👍 141 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func findLongestWord(s string, dictionary []string) string {
	var res string
	for i := range dictionary {
		if isSubString(s, dictionary[i]) {
			// 找到最长的串，并且串的位置要尽可能靠前
			if len(dictionary[i]) > len(res) || (len(dictionary[i]) == len(res) && dictionary[i] < res) {
				res = dictionary[i]
			}
		}
	}
	return res
}

func isSubString(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) {
		if s[i] == t[j] {
			j++
		}
		i++
		if j == len(t) {
			return true
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
