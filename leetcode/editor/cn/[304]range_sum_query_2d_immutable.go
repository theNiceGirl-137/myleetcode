//给定一个二维矩阵 matrix，以下类型的多个请求： 
//
// 
// 计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。 
// 
//
// 实现 NumMatrix 类： 
//
// 
// NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化 
// int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角
// (row2, col2) 的子矩阵的元素总和。 
// 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入: 
//["NumMatrix","sumRegion","sumRegion","sumRegion"]
//[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,
//1,2,2],[1,2,2,4]]
//输出: 
//[null, 8, 11, 12]
//
//解释:
//NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,
//0,1,7],[1,0,3,0,5]]]);
//numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
//numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
//numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 200 
// -105 <= matrix[i][j] <= 105 
// 0 <= row1 <= row2 < m 
// 0 <= col1 <= col2 < n 
// 最多调用 104 次 sumRegion 方法 
// 
// Related Topics 设计 数组 矩阵 前缀和 
// 👍 287 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	integral [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	// 我们可以用动态规划来计算 integral 矩阵：
	// integral[i][j] = matrix[i-1][j-1] + integral[i-1][j] + integral[i][j-1] - integral[i-1][j-1]
	// 即当前坐标的数字 + 上面长方形的数字和 + 左边长方形的数字和 - 上面长方形和左边长方形重合面积（即左上一格的长方形）中的数字和
	integral := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		integral[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			integral[i][j] = matrix[i-1][j-1]+integral[i-1][j]+integral[i][j-1]-integral[i-1][j-1]
		}
	}
	numMatrix := NumMatrix{
		integral: integral,
	}
	return numMatrix
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 即右下角与最左上角的长方形数字和 - 上面长方形的数字和 - 左边长方形的数字和 + 上面长方形和左边长方形重合面积（即左上一格的长方形）中的数字和
	return this.integral[row2+1][col2+1]-this.integral[row2+1][col1]-this.integral[row1][col2+1]+this.integral[row1][col1];
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)
