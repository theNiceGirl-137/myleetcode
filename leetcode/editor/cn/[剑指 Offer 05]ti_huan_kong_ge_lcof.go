//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。 
//
// 
//
// 示例 1： 
//
// 输入：s = "We are happy."
//输出："We%20are%20happy." 
//
// 
//
// 限制： 
//
// 0 <= s 的长度 <= 10000 
// Related Topics 字符串 👍 165 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	if s == "" {
		return ""
	}
	s1 := []byte(s)
	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' {
			n := len(s1)
			s1 = append(s1, s1[n-2:]...)
			for j := n-1; j > i+2; j-- {
				s1[j] = s1[j-2]
			}
			s1[i], s1[i+1], s1[i+2] = '%', '2', '0'
		}
	}
	return string(s1)
}
//leetcode submit region end(Prohibit modification and deletion)
