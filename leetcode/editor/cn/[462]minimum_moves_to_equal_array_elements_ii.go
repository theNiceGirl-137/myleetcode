//给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加1或减1。 您可以假设数组的长度最多为10000。 
//
// 例如: 
//
// 
//输入:
//[1,2,3]
//
//输出:
//2
//
//说明：
//只有两个动作是必要的（记得每一步仅可使其中一个元素加1或减1）： 
//
//[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
// 
// Related Topics 数组 数学 排序 
// 👍 133 👎 0


package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func minMoves2(nums []int) int {
	length, ans := len(nums), 0
	sort.Ints(nums)
	// 本质是寻找中位数，任意两个数要移动到相等一共需要移动长度差的步数，故已最中间的数作为参考会节省步数
	mid := nums[length/2]
	for i := 0; i < length; i++ {
		ans += abs(nums[i]-mid)
	}
	return ans
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
//leetcode submit region end(Prohibit modification and deletion)
