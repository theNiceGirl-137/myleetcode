//字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。 
//
// 
//
// 示例： 
//
// 
//输入：S = "ababcbacadefegdehijhklij"
//输出：[9,7,8]
//解释：
//划分结果为 "ababcbaca", "defegde", "hijhklij"。
//每个字母最多出现在一个片段中。
//像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
// 
//
// 
//
// 提示： 
//
// 
// S的长度在[1, 500]之间。 
// S只包含小写字母 'a' 到 'z' 。 
// 
// Related Topics 贪心算法 双指针 
// 👍 484 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(S string) []int {
	// 预处理一遍数组，统计字符出现的频率、第一次出现的位置、最后一次出现的位置等
	SToArray := []byte(S)
	info := make(map[byte]int)
	// 统计每个字符最后一次出现的位置
	for i, v := range SToArray {
		info[v] = i
	}
	res := make([]int, 0)
	// 当前片段开始和结束位置的下标
	start, end := 0, 0
	for i := range SToArray {
		// 一个片段结束的位置一定是结束位置最晚的某个字符的下标
		end = max763(end, info[SToArray[i]])
		if i == end {
			res = append(res, end-start+1)
			start = end+1
		}
	}
	return res
}

func max763(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
