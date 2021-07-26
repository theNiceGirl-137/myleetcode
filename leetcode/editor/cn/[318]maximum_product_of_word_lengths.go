//给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为
//每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。 
//
// 
//
// 示例 1: 
//
// 
//输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
//输出: 16 
//解释: 这两个单词为 "abcw", "xtfn"。 
//
// 示例 2: 
//
// 
//输入: ["a","ab","abc","d","cd","bcd","abcd"]
//输出: 4 
//解释: 这两个单词为 "ab", "cd"。 
//
// 示例 3: 
//
// 
//输入: ["a","aa","aaa","aaaa"]
//输出: 0 
//解释: 不存在这样的两个单词。
// 
//
// 
//
// 提示： 
//
// 
// 2 <= words.length <= 1000 
// 1 <= words[i].length <= 1000 
// words[i] 仅包含小写字母 
// 
// Related Topics 位运算 数组 字符串 
// 👍 178 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(words []string) int {
	hash := make(map[int]int)
	var ans int
	for _, word := range words {
		mask, size := 0, len(word)
		for _, ch := range word {
			// 腾出一个位置
			mask |= 1 << (ch - 'a')
		}
		// 同一字母组合选出最长的
		hash[mask] = max318(hash[mask], size)
		for h_mask, h_len := range hash {
			if mask & h_mask == 0 {
				ans = max318(ans, hash[mask]*h_len)
			}
		}
	}
	return ans
}

func max318(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
