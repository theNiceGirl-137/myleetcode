//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。 
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。 
//
// 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,0,1]
//输出：[0,1,2]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[0]
// 
//
// 示例 4： 
//
// 
//输入：nums = [1]
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// n == nums.length 
// 1 <= n <= 300 
// nums[i] 为 0、1 或 2 
// 
//
// 
//
// 进阶： 
//
// 
// 你可以不使用代码库中的排序函数来解决这道题吗？ 
// 你能想出一个仅使用常数空间的一趟扫描算法吗？ 
// 
// Related Topics 排序 数组 双指针 
// 👍 865 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	// 分别设置两个索引 zero 和 two，保证下标 0 到 zero 对应的数组元素都为 0，下标 two 到 numsSize-1 对应的数组元素都是 2
	zero, two := -1, len(nums)
	for i := 0; i < two; {
		// 直接将该元素加入到数组中存 1 的部分
		if nums[i] == 1 {
			i++
		// 交换 zero 前面一个元素 1，将此时的 0 放入 zero 数组
		// i 与 zero 都右移
		} else if nums[i] == 0 {
			zero++
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		// 直接将这个元素加入 two 数组部分
		} else {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
