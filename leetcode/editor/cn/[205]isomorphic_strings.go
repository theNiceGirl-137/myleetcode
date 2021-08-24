//给定两个字符串 s 和 t，判断它们是否是同构的。 
//
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。 
//
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。 
//
// 
//
// 示例 1: 
//
// 
//输入：s = "egg", t = "add"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "foo", t = "bar"
//输出：false 
//
// 示例 3： 
//
// 
//输入：s = "paper", t = "title"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 可以假设 s 和 t 长度相同。 
// 
// Related Topics 哈希表 字符串 
// 👍 379 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	// 记录两个字符串每个位置的字符第一次出现的位置，如果两个字符串中相同位置的字符与它们第一次出现的位置一样，那么这两个字符串同构
	sFirst := make([]int, 256)
	tFirst := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if sFirst[s[i]] != tFirst[t[i]] {
			return false
		}
		// 初始化为 0，需要加一来排除相同位置字符从下标 0 开始的情况
		sFirst[s[i]] = i+1
		tFirst[t[i]] = i+1
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
