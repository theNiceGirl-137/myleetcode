//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。 
//
// 
//
// 示例 1： 
//
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "AB
//CCED"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SE
//E"
//输出：true
// 
//
// 示例 3： 
//
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "AB
//CB"
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// m == board.length 
// n = board[i].length 
// 1 <= m, n <= 6 
// 1 <= word.length <= 15 
// board 和 word 仅由大小写英文字母组成 
// 
//
// 
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？ 
// Related Topics 数组 回溯算法 
// 👍 892 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	find := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backTracking79(i, j, board, word, &find, visited, 0)
		}
	}
	return find
}

func backTracking79(i, j int, board [][]byte, word string, find *bool, visited [][]bool, pos int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if visited[i][j] == true || *find == true || board[i][j] != word[pos] {
		return
	}
	if pos == len(word)-1 {
		*find = true
		return
	}
	visited[i][j] = true
	backTracking79(i+1, j, board, word, find, visited, pos+1)
	backTracking79(i-1, j, board, word, find, visited, pos+1)
	backTracking79(i, j+1, board, word, find, visited, pos+1)
	backTracking79(i, j-1, board, word, find, visited, pos+1)
	visited[i][j] = false
}
//leetcode submit region end(Prohibit modification and deletion)
