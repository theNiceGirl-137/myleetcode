//假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。 
//
// 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则
//的情况下种入 n 朵花？能则返回 true ，不能则返回 false。 
//
// 
//
// 示例 1： 
//
// 
//输入：flowerbed = [1,0,0,0,1], n = 1
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：flowerbed = [1,0,0,0,1], n = 2
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= flowerbed.length <= 2 * 104 
// flowerbed[i] 为 0 或 1 
// flowerbed 中不存在相邻的两朵花 
// 0 <= n <= flowerbed.length 
// 
// Related Topics 贪心算法 数组 
// 👍 339 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	// 很容易分析得知连续三个零可以种一朵花
	// 为了方便处理边界情况，可以在两边补 0，假设为没种花的情况
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	// 从 1 开始遍历，遍历到倒数第二位
	for i := 1; i < len(flowerbed)-1; i++ {
		// 连续三个 0
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
		if n < 1 {
			return true
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
