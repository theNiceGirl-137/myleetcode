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

// https://leetcode-cn.com/problems/shortest-bridge/solution/go-by-louv-0bbj/
func shortestBridge(A [][]int) int {
	queue := make([][2]int, 0)
	flipped := false
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				dfs934(A, &queue, i, j)
				flipped = true
				break
			}
		}
		if flipped {
			break
		}
	}
	level := 0
	for len(queue) > 0 {
		level++
		for i := 0; i < len(queue); i++ {
			temp := queue[0]
			queue = queue[1:]
			for j := 0; j < 4; j++ {
				x, y := temp[0]+direction934[j], temp[1]+direction934[j+1]
				if x >= 0 && y >= 0 && x < len(A) && y < len(A[0]) && A[x][y] >= 0 {
					if A[x][y] == 1 {
						return level
					}
					A[x][y] = -1
					queue = append(queue, [2]int{x, y})
				}
			}
		}
	}
	return level
}

func dfs934(A [][]int, queue *[][2]int, i, j int) {
	A[i][j] = -1
	for k := 0; k < 4 ; k++ {
		x, y := i+direction934[k], j+direction934[k+1]
		if x >= 0 && y >= 0 && x < len(A) && y < len(A[0]) {
			if A[x][y] == 0 {
				*queue = append(*queue, [2]int{x, y})
			}
			if A[x][y] == 1 {
				dfs934(A, queue, x, y)
			}
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
