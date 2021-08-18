//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。 
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。 
//
// 示例 2： 
//
// 
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 105 
// -109 <= nums[i] <= 109 
// 
// Related Topics 并查集 数组 哈希表 
// 👍 878 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	var ans int
	if len(nums) < 1 {
		return ans
	}
	mapper := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		mapper[nums[i]] = true
	}
	ans = 1
	for num := range mapper {
		if mapper[num-1] {
			continue
		}
		current := 1
		for mapper[num+1] {
			num++
			current++
		}
		if current > ans {
			ans = current
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
