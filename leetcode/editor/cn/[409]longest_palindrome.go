//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。 
//
// 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。 
//
// 注意: 
//假设字符串的长度不会超过 1010。 
//
// 示例 1: 
//
// 
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
// 
// Related Topics 贪心 哈希表 字符串 👍 328 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome409(s string) int {
	var count int
	mapper := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		mapper[s[i]]++
	}
	for _, value := range mapper {
		if value/2 >= 1 {
			count += value/2*2
		}
	}
	// 是否添加一个单独字符放在回文串中间位置
	if len(s)-count > 0 {
		count += 1
	}
	return count
}
//leetcode submit region end(Prohibit modification and deletion)
