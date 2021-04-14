//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。 
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下： 
//
// 
//P   A   H   N
//A P L S I I G
//Y   I   R 
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。 
//
// 请你实现这个将字符串进行指定行数变换的函数： 
//
// 
//string convert(string s, int numRows); 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
// 
//示例 2：
//
// 
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
// 
//
// 示例 3： 
//
// 
//输入：s = "A", numRows = 1
//输出："A"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 由英文字母（小写和大写）、',' 和 '.' 组成 
// 1 <= numRows <= 1000 
// 
// Related Topics 字符串 
// 👍 1109 👎 0


package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}
	b := []rune(s)
	// 构成 Z 字型的每一行的字符串
	res := make([]string, numRows)
	// 观察规律可以知道，可以把 Z 字型分解为 V 字型
	// 每一个 V 字是 2*n-2 个字符，按照每组 V 字依次填充
	period := 2*numRows - 2
	for i := range b {
		mod := i % period
		if mod < numRows {
			res[mod] += string(b[i])
		} else {
			// 如果超出numRows需要往回走
			res[period-mod] += string(b[i])
		}
	}
	return strings.Join(res, "")
}
//leetcode submit region end(Prohibit modification and deletion)
