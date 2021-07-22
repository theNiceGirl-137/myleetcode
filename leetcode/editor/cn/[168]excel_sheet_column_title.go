//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。 
//
// 例如： 
//
// 
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28 
//...
// 
//
// 
//
// 示例 1： 
//
// 
//输入：columnNumber = 1
//输出："A"
// 
//
// 示例 2： 
//
// 
//输入：columnNumber = 28
//输出："AB"
// 
//
// 示例 3： 
//
// 
//输入：columnNumber = 701
//输出："ZY"
// 
//
// 示例 4： 
//
// 
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= columnNumber <= 231 - 1 
// 
// Related Topics 数学 字符串 
// 👍 434 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	// 本题本质就是进制转换，10 进制转 26 进制，但有所不同的是正常转换成26进制的余数是 0-25
	// 而本题的余数是 1-26 (对应A-Z)，为了消除差距的这个1，有两种方法：
	// 一 让除数减一，那么余数自然就少一，原来余 1 的变成余 0，当 n 为 26 时，n-1 的余数变为 25，再加 'A' 可得 'Z'
	// 二 对值为 26 的倍数单独处理，保留为 26，而不取余
	var res []byte
	for columnNumber > 0 {
		columnNumber--
		res = append([]byte{byte(columnNumber%26+65)}, res...)
		columnNumber /=26
	}
	return string(res)
}
//leetcode submit region end(Prohibit modification and deletion)
