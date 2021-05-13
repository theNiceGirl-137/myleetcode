//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。 
//
// candidates 中的每个数字在每个组合中只能使用一次。 
//
// 说明： 
//
// 
// 所有数字（包括目标数）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
// 
//
// 示例 2: 
//
// 输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//] 
// Related Topics 数组 回溯算法 
// 👍 575 👎 0


package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	if len(candidates) < 1 {
		return ans
	}
	sort.Ints(candidates)
	comb := make([]int, 0)
	backTracking40(&ans, comb, candidates, target, 0, 0)
	return ans
}

func backTracking40(ans *[][]int, comb, candidate []int, target, sum,  pos int) {
	if sum >= target {
		if sum == target {
			temp := make([]int, len(comb))
			copy(temp, comb)
			*ans = append(*ans, temp)
		}
		return
	}
	for i := pos; i < len(candidate); i++ {
		if i-1 >= pos && candidate[i-1] == candidate[i] {
			continue
		}
		comb = append(comb, candidate[i])
		backTracking40(ans, comb, candidate, target, sum+candidate[i], i+1)
		comb = comb[:len(comb)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
