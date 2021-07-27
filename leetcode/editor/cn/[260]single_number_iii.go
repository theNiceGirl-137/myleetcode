//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。 
//
// 
//
// 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？ 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,1,3,2,5]
//输出：[3,5]
//解释：[5, 3] 也是有效的答案。
// 
//
// 示例 2： 
//
// 
//输入：nums = [-1,0]
//输出：[-1,0]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0,1]
//输出：[1,0]
// 
//
// 提示： 
//
// 
// 2 <= nums.length <= 3 * 104 
// -231 <= nums[i] <= 231 - 1 
// 除两个只出现一次的整数外，nums 中的其他数字都出现两次 
// 
// Related Topics 位运算 数组 
// 👍 420 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber260(nums []int) []int {
	// a ^ a = 0
	// ans一定不为 0，因为一定是两个不一样的数异或
	var temp, bitIndex int
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		temp ^= nums[i]
	}
	// 不一样的数异或至少有一位为 1，因为
	// 1 ^ 1 = 0
	// 1 ^ 0 = 1
	// 0 ^ 1 = 1
	// 0 ^ 0 = 0
	// 故根据 1 的位置不同分为两组，必然会将两个数分开
	bitIndex = 1
	for {
		if temp & 1 == 1 {
			break
		}
		temp >>= 1
		bitIndex <<= 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] & bitIndex == 0 {
			res[0] ^= nums[i]
		} else {
			res[1] ^= nums[i]
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
