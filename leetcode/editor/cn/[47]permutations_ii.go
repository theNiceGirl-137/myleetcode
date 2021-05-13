//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 8 
// -10 <= nums[i] <= 10 
// 
// Related Topics 回溯算法 
// 👍 694 👎 0


package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	backTracking47(path, nums, used, &ans)
	return ans
}

func backTracking47(path, nums []int, used []bool, ans *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 在有重复数组的情况下需要剪枝去重
		// 如果与前一个数相同，说明前一个数已经选择完了可能的组合，直接跳过
		if i-1 >= 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		// 如果当前数字已经被选择过，也直接跳过
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backTracking47(path, nums, used, ans)
		path = path[:len(path)-1]
		used[i] = false
	}
}
//leetcode submit region end(Prohibit modification and deletion)
