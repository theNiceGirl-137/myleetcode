//给定一个包含了一些 0 和 1 的非空二维数组 grid 。 
//
// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 
//0（代表水）包围着。 
//
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。) 
//
// 
//
// 示例 1: 
//
// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,1,1,0,1,0,0,0,0,0,0,0,0],
// [0,1,0,0,1,1,0,0,1,0,1,0,0],
// [0,1,0,0,1,1,0,0,1,1,1,0,0],
// [0,0,0,0,0,0,0,0,0,0,1,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,0,0,0,0,0,0,1,1,0,0,0,0]]
// 
//
// 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。 
//
// 示例 2: 
//
// [[0,0,0,0,0,0,0,0]] 
//
// 对于上面这个给定的矩阵, 返回 0。 
//
// 
//
// 注意: 给定的矩阵grid 的长度和宽度都不超过 50。 
// Related Topics 深度优先搜索 数组 
// 👍 482 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maxAreaOfIsland(grid [][]int) int {
	//return stackWay(grid)
	//return recursionWay1(grid)
	return recursionWay2(grid)
}

// 栈
// 对于四个方向的遍历，每相邻两位即为上下左右四个方向之一
var direction695 = []int{-1, 0, 1, 0, -1}

func stackWay(grid [][]int) int {
	var localArea, area, x, y int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				localArea = 1
				grid[i][j] = 0
				island := make([][2]int, 0)
				island = append(island, [2]int{i, j})
				for len(island) != 0 {
					top := island[len(island)-1]
					island = island[:len(island)-1]
					// 遍历相连的四个点
					for k := 0; k < 4; k++ {
						x = top[0]+direction695[k]
						y = top[1]+direction695[k+1]
						if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
							grid[x][y] = 0
							localArea++
							island = append(island, [2]int{x, y})
						}
					}
				}
				area = max695(area, localArea)
			}
		}
	}
	return area
}

// 递归写法一
func recursionWay1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max695(maxArea, dfs695(grid, i, j))
			}
		}
	}
	return maxArea
}

func dfs695(grid [][]int, r, c int) int {
	if grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	var x, y int
	area := 1
	for i := 0; i < 4; i++ {
		x = r + direction695[i]
		y = c + direction695[i+1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			area += dfs695(grid, x, y)
		}
	}
	return area
}

// 递归写法二
func recursionWay2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maxArea = max695(maxArea, dfs695_2(grid, i, j))
		}
	}
	return maxArea
}

func dfs695_2(grid [][]int, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	return 1 + dfs695_2(grid, r+1, c) + dfs695_2(grid, r-1, c) + dfs695_2(grid, r, c+1) + dfs695_2(grid, r, c-1)
}

func max695(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
