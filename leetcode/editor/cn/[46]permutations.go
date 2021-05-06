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
	backTracking46(nums, 0, &ans)
	return ans
}

func backTracking46(nums []int, level int, ans *[][]int) {
	if level == len(nums)-1 {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*ans = append(*ans, temp)
		return
	}
	for  i := level; i < len(nums); i++{
		nums[i], nums[level] = nums[level], nums[i]
		backTracking46(nums, level+1, ans)
		nums[i], nums[level] = nums[level], nums[i]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
