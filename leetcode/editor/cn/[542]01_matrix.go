//给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。 
//
// 两个相邻元素间的距离为 1 。 
//
// 
//
// 示例 1： 
//
// 
//输入：
//[[0,0,0],
// [0,1,0],
// [0,0,0]]
//
//输出：
//[[0,0,0],
// [0,1,0],
// [0,0,0]]
// 
//
// 示例 2： 
//
// 
//输入：
//[[0,0,0],
// [0,1,0],
// [1,1,1]]
//
//输出：
//[[0,0,0],
// [0,1,0],
// [1,2,1]]
// 
//
// 
//
// 提示： 
//
// 
// 给定矩阵的元素个数不超过 10000。 
// 给定矩阵中至少有一个元素是 0。 
// 矩阵中的元素只在四个方向上相邻: 上、下、左、右。 
// 
// Related Topics 深度优先搜索 广度优先搜索 
// 👍 431 👎 0


package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func updateMatrix(mat [][]int) [][]int {
	// 一般来说，这道题涉及到四个方向上的最近搜索，使用 bfs 可能会使时间复杂度达到 O(m^2n^2)
	// 一种方法是使用一个 dp 数组做空间压缩，使得 bfs 不会遍历重复的位置
	// 另一种方法是从左上到右下进行一次动态搜索，再从右下到左上进行一次动态搜索，两次动态搜索完成四个方向的查找
	if len(mat) < 1 {
		return nil
	}
	// dp 数组默认初始化所有值都为最大值
	dp := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		dp[i] = make([]int, len(mat[0]))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// 从左上到右下搜索，只关注右和下两个方向
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
			}
		}
	}
	// 从右下到左上搜索，只关注左和上两个方向
	for i := len(mat)-1; i >= 0; i-- {
		for j := len(mat[0])-1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if j < len(mat[0])-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
				if i < len(mat)-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
			}
		}
	}
	return dp
}
//leetcode submit region end(Prohibit modification and deletion)
