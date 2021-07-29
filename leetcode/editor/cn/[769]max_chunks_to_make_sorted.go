//数组arr是[0, 1, ..., arr.length - 1]的一种排列，我们将这个数组分割成几个“块”，并将这些块分别进行排序。之后再连接起来，使得连
//接的结果和按升序排序后的原数组相同。 
//
// 我们最多能将数组分成多少块？ 
//
// 示例 1: 
//
// 输入: arr = [4,3,2,1,0]
//输出: 1
//解释:
//将数组分成2块或者更多块，都无法得到所需的结果。
//例如，分成 [4, 3], [2, 1, 0] 的结果是 [3, 4, 0, 1, 2]，这不是有序的数组。
// 
//
// 示例 2: 
//
// 输入: arr = [1,0,2,3,4]
//输出: 4
//解释:
//我们可以把它分成两块，例如 [1, 0], [2, 3, 4]。
//然而，分成 [1, 0], [2], [3], [4] 可以得到最多的块数。
// 
//
// 注意: 
//
// 
// arr 的长度在 [1, 10] 之间。 
// arr[i]是 [0, 1, ..., arr.length - 1]的一种排列。 
// 
// Related Topics 栈 贪心 数组 排序 单调栈 
// 👍 142 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maxChunksToSorted(arr []int) int {
	// 从左往右遍历，同时记录当前的最大值，每当当前最大值等于数组位置时，我们可以多一次分割
	// 为什么可以通过这个算法解决问题呢？如果当前最大值大于数组位置，则说明右边一定有小
	// 于数组位置的数字，需要把它也加入待排序的子数组；又因为数组只包含不重复的 0 到 n，所以
	// 当前最大值一定不会小于数组位置
	// 所以每当当前最大值等于数组位置时，假设为 p，我们可以成功完成一次分割，并且其与上一次分割位置 q 之间的值一定是 q + 1 到 p 的所有数字
	chunks, maxCur := 0, 0
	for i := 0; i < len(arr); i++ {
		maxCur = max(maxCur, arr[i])
		if maxCur == i {
			chunks++
		}
	}
	return chunks
}
//leetcode submit region end(Prohibit modification and deletion)
