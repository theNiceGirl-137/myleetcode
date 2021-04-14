//编写一个函数来查找字符串数组中的最长公共前缀。 
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["flower","flow","flight"]
//输出："fl"
// 
//
// 示例 2： 
//
// 
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 
//
// 
//
// 提示： 
//
// 
// 0 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] 仅由小写英文字母组成 
// 
// Related Topics 字符串 
// 👍 1562 👎 0


package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// 作为前缀的子串一定在第 0 位
		if strings.Index(strs[i], prefix) != 0 {
			for prefix != "" {
				if strings.Index(strs[i], prefix) == 0 {
					break
				}
				prefix = prefix[:len(prefix)-1]
			}
		}
	}
	return prefix
}
//leetcode submit region end(Prohibit modification and deletion)
