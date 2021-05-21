//编写一个程序，通过填充空格来解决数独问题。 
//
// 数独的解法需 遵循如下规则： 
//
// 
// 数字 1-9 在每一行只能出现一次。 
// 数字 1-9 在每一列只能出现一次。 
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图） 
// 
//
// 数独部分空格内已填入了数字，空白格用 '.' 表示。 
//
// 
//
// 
// 
// 
// 示例： 
//
// 
//输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5","."
//,".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".","."
//,"3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"
//],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],["
//.",".",".",".","8",".",".","7","9"]]
//输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"
//],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["
//4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","
//6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","
//5","2","8","6","1","7","9"]]
//解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
//
//
// 
//
// 
//
// 提示： 
//
// 
// board.length == 9 
// board[i].length == 9 
// board[i][j] 是一位数字或者 '.' 
// 题目数据 保证 输入数独仅有一个解 
// 
// 
// 
// 
// Related Topics 哈希表 回溯算法 
// 👍 836 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte)  {
	rows := make([]map[int]bool, 9)
	columns := make([]map[int]bool, 9)
	blocks := make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]bool, 9)
		columns[i] = make(map[int]bool, 9)
		blocks[i] = make(map[int]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				v := int(board[i][j]-'0')
				rows[i][v] = true
				columns[j][v] = true
				blocks[findBlock(i, j)][v] = true
			}
		}
	}
	backTracking37(0, 0, board, rows, columns, blocks)
}

func backTracking37(a, b int, board [][]byte, rows, columns, blocks []map[int]bool) bool {
	if a > 8 {
		a = 0
		b++
	}
	if b > 8 {
		return true
	}
	if board[a][b] != '.' {
		return backTracking37(a+1, b, board, rows, columns, blocks)
	}
	for i := 1; i <= 9; i++ {
		if rows[a][i] || columns[b][i] || blocks[findBlock(a, b)][i] {
			continue
		}
		board[a][b] = byte(i)+'0'
		rows[a][i], columns[b][i], blocks[findBlock(a, b)][i] = true, true, true
		if backTracking37(a+1, b, board, rows, columns, blocks)  {
			return true
		}
		board[a][b] = '.'
		rows[a][i], columns[b][i], blocks[findBlock(a, b)][i] = false, false, false
	}
	return false
}

func findBlock(i, j int) int {
	return i/3 + j/3*3
}
//leetcode submit region end(Prohibit modification and deletion)
