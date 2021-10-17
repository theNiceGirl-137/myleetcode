//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。 
//
// 
//
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。 
//
// 
//
// 
//
// 示例 1： 
//
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = 
//"ABCCED"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= board.length <= 200 
// 1 <= board[i].length <= 200 
// board 和 word 仅由大小写英文字母组成 
// 
//
// 
//
// 注意：本题与主站 79 题相同：https://leetcode-cn.com/problems/word-search/ 
// Related Topics 数组 回溯 矩阵 👍 429 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	var find bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backTrackingOffer12(board, visited, word, i, j, 0, &find)
		}
	}
	return find
}

func backTrackingOffer12(board [][]byte, visited [][]bool, word string, i, j, pos int, find *bool) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || pos >= len(word) {
		return
	}
	if visited[i][j] || *find || board[i][j] != word[pos] {
		return
	}
	if len(word)-1 == pos {
		*find = true
		return
	}
	visited[i][j] = true
	backTrackingOffer12(board, visited, word, i+1, j, pos+1, find)
	backTrackingOffer12(board, visited, word, i-1, j, pos+1, find)
	backTrackingOffer12(board, visited, word, i, j+1, pos+1, find)
	backTrackingOffer12(board, visited, word, i, j-1, pos+1, find)
	visited[i][j] = false
}
//leetcode submit region end(Prohibit modification and deletion)
