//给定一个字符串，请将字符串里的字符按照出现的频率降序排列。 
//
// 示例 1: 
//
// 
//输入:
//"tree"
//
//输出:
//"eert"
//
//解释:
//'e'出现两次，'r'和't'都只出现一次。
//因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
// 
//
// 示例 2: 
//
// 
//输入:
//"cccaaa"
//
//输出:
//"cccaaa"
//
//解释:
//'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
//注意"cacaca"是不正确的，因为相同的字母必须放在一起。
// 
//
// 示例 3: 
//
// 
//输入:
//"Aabb"
//
//输出:
//"bbAa"
//
//解释:
//此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
//注意'A'和'a'被认为是两种不同的字符。
// 
// Related Topics 堆 哈希表 
// 👍 237 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(s string) string {
	counts := make(map[byte]int)
	maxCount := 0
	// 统计每个字符的出现次数
	for i := range s {
		counts[s[i]]++
		maxCount = max(maxCount, counts[s[i]])
	}
	buckets := make([][]byte, maxCount+1)
	res := make([]byte, 0)
	// 按出现的次数将字符放进新桶内
	// 出现次数越大，即 v 值更靠后，就是出现频率更高的字符
	for k, v := range counts {
		buckets[v] = append(buckets[v], k)
	}
	for i := len(buckets)-1; i > 0; i-- {
		if len(buckets[i]) > 0 {
			for j := range buckets[i] {
				// 按出现的次数追加字符
				for k := 0; k < counts[buckets[i][j]]; k++ {
					res = append(res, buckets[i][j])
				}
			}
		}
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
