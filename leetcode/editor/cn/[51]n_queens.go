//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。 
//
// 
// 
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：[["Q"]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 9 
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。 
// 
// 
// 
// Related Topics 回溯算法 
// 👍 862 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var ans [][]string
	if n == 0 {
		return ans
	}
	var str string
	for i := 0; i < n; i++ {
		str += "."
	}
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = str
	}
	column := make([]bool, n)
	ldiag := make([]bool, 2*n)
	rdiag := make([]bool, 2*n)
	backTracking51(&ans, board, column, ldiag, rdiag, 0, n)
	return ans
}

func backTracking51(ans *[][]string, board []string, column, ldiag, rdiag []bool, row, n int) {
	if row == n {
		temp := make([]string, len(board))
		copy(temp, board)
		*ans = append(*ans, temp)
		return
	}
	// 每个 Q 的行、列、正反对角线都不能相等
	// 在每一行遍历，确定行不相等，将确定的列值保存下来，确定列不相等
	// 可以发现，同一正对角线的行与列的差值是相等的，可以用一维数组存储，直接表示一条对角线
	// 同理，同一反对角线的行与列的和是相等的，也可以用一维数组存储，直接表示一条对角线
	for i := 0; i < n; i++ {
		if column[i] || ldiag[row-i+n] || rdiag[row+i] {
			continue
		}
		c := []rune(board[row])
		c[i] = 'Q'
		board[row] = string(c)
		column[i], ldiag[row-i+n], rdiag[row+i] = true, true, true
		backTracking51(ans, board, column, ldiag, rdiag, row+1, n)
		c[i] = '.'
		board[row] = string(c)
		column[i], ldiag[row-i+n], rdiag[row+i] = false, false, false
	}
}
//leetcode submit region end(Prohibit modification and deletion)
