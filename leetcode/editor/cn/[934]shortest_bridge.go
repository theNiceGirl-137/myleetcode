//在给定的二维二进制数组 A 中，存在两座岛。（岛是由四面相连的 1 形成的一个最大组。） 
//
// 现在，我们可以将 0 变为 1，以使两座岛连接起来，变成一座岛。 
//
// 返回必须翻转的 0 的最小数目。（可以保证答案至少是 1 。） 
//
// 
//
// 示例 1： 
//
// 
//输入：A = [[0,1],[1,0]]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：A = [[0,1,0],[0,0,0],[0,0,1]]
//输出：2
// 
//
// 示例 3： 
//
// 
//输入：A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//输出：1 
//
// 
//
// 提示： 
//
// 
// 2 <= A.length == A[0].length <= 100 
// A[i][j] == 0 或 A[i][j] == 1 
// 
// Related Topics 深度优先搜索 广度优先搜索 
// 👍 153 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
var direction934 = []int{-1, 0, 1, 0, -1}

func shortestBridge(A [][]int) int {
	m, n := len(A), len(A[0])
	points := make([][2]int, 0)
	flipped := false
	// dfs 寻找第一个岛屿，并把 1 全部赋值为 2
	for i := 0; i < m; i++ {
		if flipped {
			break
		}
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				dfs934(&points, A, m, n, i, j)
				flipped = true
				break
			}
		}
	}
	// bfs 寻找第二个岛屿，并把过程中经过的 0 赋值为 2
	var x, y, level int
	for len(points) > 0 {
		level++
		nPoints := len(points)
		for nPoints > 0 {
			temp := points[0]
			points = points[1:]
			nPoints--
			for k := 0; k < 4; k++ {
				x = temp[0]+direction934[k]
				y = temp[1]+direction934[k+1]
				if x >= 0 && y >= 0 && x <m && y < n {
					if A[x][y] == 2 {
						continue
					}
					if A[x][y] ==  1 {
						return level
					}
					points = append(points, [2]int{x, y})
					A[x][y] = 2
				}
			}
		}
	}
	return 0
}

func dfs934(points *[][2]int, A [][]int, m, n, i, j int) {
	if i < 0 || j < 0 || i == m || j == n || A[i][j] == 2 {
		return
	}
	if A[i][j] == 0 {
		*points = append(*points, [2]int{i, j})
		return
	}
	A[i][j] = 2
	dfs934(points, A, m, n, i-1, j)
	dfs934(points, A, m, n, i+1, j)
	dfs934(points, A, m, n, i, j-1)
	dfs934(points, A, m, n, i, j+1)
}
//leetcode submit region end(Prohibit modification and deletion)
