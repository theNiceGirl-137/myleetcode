//给定一个 没有重复 数字的序列，返回其所有可能的全排列。 
//
// 示例: 
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//] 
// Related Topics 回溯算法 
// 👍 1317 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	var ans [][]int
	var backTracking func(path []int)
	visited := make(map[int]bool)
	backTracking = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for _, v := range nums {
			if visited[v] {
				continue
			}
			path = append(path, v)
			visited[v] = true
			backTracking(path)
			path = path[:len(path)-1]
			visited[v] = false
		}
	}
	backTracking([]int{})
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
