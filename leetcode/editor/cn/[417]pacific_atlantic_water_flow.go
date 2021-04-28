//给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。 
//
// 规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。 
//
// 请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。 
//
// 
//
// 提示： 
//
// 
// 输出坐标的顺序不重要 
// m 和 n 都小于150 
// 
//
// 
//
// 示例： 
//
// 
//
// 
//给定下面的 5x5 矩阵:
//
//  太平洋 ~   ~   ~   ~   ~ 
//       ~  1   2   2   3  (5) *
//       ~  3   2   3  (4) (4) *
//       ~  2   4  (5)  3   1  *
//       ~ (6) (7)  1   4   5  *
//       ~ (5)  1   1   2   4  *
//          *   *   *   *   * 大西洋
//
//返回:
//
//[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).
// 
//
// 
// Related Topics 深度优先搜索 广度优先搜索 
// 👍 232 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
var direction417 = []int{-1, 0, 1, 0, -1}

// 如果对所有位置进行搜索，在不剪枝的情况下复杂度会很高，因此可以反过来想，从两个大洋
// 往上流，只需要对矩形四条边进行搜索
func pacificAtlantic(heights [][]int) [][]int {
	ans := make([][]int, 0)
	if len(heights) == 0 || len(heights[0]) == 0 {
		return ans
	}
	m, n := len(heights), len(heights[0])
	canReachP := make([][]bool, m)
	canReachA := make([][]bool, m)
	for i := 0; i < m; i++ {
		canReachP[i] = make([]bool, n)
		canReachA[i] = make([]bool, n)
	}
	// 分别搜索四条边
	for i := 0; i < m; i++ {
		dfs417(heights, canReachP, i, 0)
		dfs417(heights, canReachA, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs417(heights, canReachP, 0, i)
		dfs417(heights, canReachA, m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachP[i][j] && canReachA[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func dfs417(heights [][]int, canReach [][]bool, r, c int) {
	if canReach[r][c] {
		return
	}
	canReach[r][c] = true
	// 遍历周围的结点
	for i := 0; i < 4; i++ {
		x := r + direction417[i]
		y := c + direction417[i+1]
		if x >= 0 && x < len(heights) && y >= 0 && y < len(heights[0]) && heights[r][c] <= heights[x][y] {
			dfs417(heights, canReach, x, y)
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
