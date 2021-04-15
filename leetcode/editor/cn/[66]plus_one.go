//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。 
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。 
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。 
//
// 
//
// 示例 1： 
//
// 
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
// 
//
// 示例 2： 
//
// 
//输入：digits = [4,3,2,1]
//输出：[4,3,2,2]
//解释：输入数组表示数字 4321。
// 
//
// 示例 3： 
//
// 
//输入：digits = [0]
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= digits.length <= 100 
// 0 <= digits[i] <= 9 
// 
// Related Topics 数组 
// 👍 667 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func plusOne(digits []int) []int {
	// 标记是否有进位
	isOverAdd := 0
	digits[len(digits)-1]++
	for i := len(digits)-1; i >= 0; i-- {
		digits[i] += isOverAdd
		isOverAdd = 0
		if digits[i] >= 10 {
			isOverAdd = 1
			digits[i] %= 10
		}
	}
	if isOverAdd == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
//leetcode submit region end(Prohibit modification and deletion)
