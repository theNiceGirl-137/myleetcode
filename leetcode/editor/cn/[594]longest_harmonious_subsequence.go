//和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。 
//
// 现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。 
//
// 数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,2,2,5,2,3,7]
//输出：5
//解释：最长的和谐子序列是 [3,2,2,2,3]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3,4]
//输出：2
// 
//
// 示例 3： 
//
// 
//输入：nums = [1,1,1,1]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2 * 104 
// -109 <= nums[i] <= 109 
// 
// Related Topics 数组 哈希表 排序 
// 👍 182 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func findLHS(nums []int) int {
	var ans int
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]]++
	}
	for key, value := range hashmap {
		if hashmap[key+1] != 0 {
			ans = max(ans, value+hashmap[key+1])
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
