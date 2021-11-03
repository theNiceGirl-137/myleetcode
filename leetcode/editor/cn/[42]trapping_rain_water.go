//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
// 
//
// 示例 2： 
//
// 
//输入：height = [4,2,0,3,2,5]
//输出：9
// 
//
// 
//
// 提示： 
//
// 
// n == height.length 
// 1 <= n <= 2 * 10⁴ 
// 0 <= height[i] <= 10⁵ 
// 
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 2821 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	sum := 0
	stack := make([]int, 0)
	current := 0
	for current < len(height) {
		// 如果栈不空并且当前指向的高度大于栈顶高度就一直循环
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			// 取出要出栈的元素
			h := height[stack[len(stack)-1]]
			// 出栈
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			// 两堵墙之前的距离
			distance := current-stack[len(stack)-1]-1
			// 不确定柱子的前后高矮顺序，选择较矮的一个
			min := min(height[stack[len(stack)-1]], height[current])
			sum += distance*(min-h)
		}
		// 当前指向的墙入栈
		stack = append(stack, current)
		current++
	}
	return sum
}
//leetcode submit region end(Prohibit modification and deletion)
