//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
//。
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X"
//,"X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
//会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
// 
//
// 示例 2： 
//
// 
//输入：board = [["X"]]
//输出：[["X"]]
// 
//
// 
//
// 提示： 
//
// 
// m == board.length 
// n == board[i].length 
// 1 <= m, n <= 200 
// board[i][j] 为 'X' 或 'O' 
// 
// 
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集 
// 👍 533 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func solve(board [][]byte) {
	if len(board) < 1 {
		return
	}
	// 将边界上与及其相连的 O 全部标记为 #
	// 剩下未标记的 O 就是被 X 包围的
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if (i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1) && board[i][j] == 'O' {
				dfs130(board, i, j)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs130(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	dfs130(board, i+1, j)
	dfs130(board, i-1, j)
	dfs130(board, i, j+1)
	dfs130(board, i, j-1)
}
//leetcode submit region end(Prohibit modification and deletion)
