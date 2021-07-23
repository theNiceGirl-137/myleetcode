//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。 
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：[3,2,3]
//输出：3 
//
// 示例 2： 
//
// 
//输入：[2,2,1,1,1,2,2]
//输出：2
// 
//
// 
//
// 进阶： 
//
// 
// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。 
// 
// Related Topics 数组 哈希表 分治 计数 排序 
// 👍 1075 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// Boyer–Moore majority vote algorithm 摩尔投票算法
	// 删去一個数列中的两个不同的数字，不会影响该数列的 majority element
	// 假想有一群人要投票，候选人有 A、B、C，假设 A 已知会过半数的话
	// 一 被取消资格的是 B 跟 C -> 显然 A 还是过半数
	// 二 被取消资格的是 (A, B) 或 (A, C) -> 一个投 A 的和一个不投 A 的同步被排除
	// 同理，在不只 3 个候选人(数字)的时候，每次取 2 个人取消投票资格(移除)，亦无法改变投票结果( A 还是会是majority element)
	res, cnt := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			cnt++
		} else if cnt > 0 {
			// 如果不相同，将这个数与上一部分数一起移除 --> 对应情况一与二
			cnt--
		} else {
			res = nums[i]
		}
	}
	return res
	// 或者在已知数组中会有一个数占多半，那在排序之后这个数一定在数组最中间
	//sort.Ints(nums)
	//return nums[len(nums)/2]
}
//leetcode submit region end(Prohibit modification and deletion)
