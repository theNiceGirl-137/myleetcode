//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性： 
//
// 
// 每行的元素从左到右升序排列。 
// 每列的元素从上到下升序排列。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 5
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 20
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= n, m <= 300 
// -109 <= matix[i][j] <= 109 
// 每行的所有元素从左到右升序排列 
// 每列的所有元素从上到下升序排列 
// -109 <= target <= 109 
// 
// Related Topics 数组 二分查找 分治 矩阵 
// 👍 676 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	// 这道题有一个简单的技巧：我们可以从右上角开始查找，若当前值大于待搜索值，我们向左
	// 移动一位；若当前值小于待搜索值，我们向下移动一位
	// 如果最终移动到左下角时仍不等于待搜索值，则说明待搜索值不存在于矩阵中
	m, n := len(matrix), len(matrix[0])
	if m < 1 {
		return false
	}
	i, j := 0, n-1
	for i < m && j >= 0 {
		temp := matrix[i][j]
		if temp == target {
			return true
		} else if temp > target {
			j--
		} else {
			i++
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
