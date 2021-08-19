//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。 
//
// 
//
// 提示： 
//
// 
// num1 和num2 的长度都小于 5100 
// num1 和num2 都只包含数字 0-9 
// num1 和num2 都不包含任何前导零 
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式 
// 
// Related Topics 数学 字符串 模拟 
// 👍 403 👎 0


package leetcode

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	var output string
	var addBit uint8
	num1, num2 = reverseStr(num1), reverseStr(num2)
	oneLen, twoLen := len(num1), len(num2)
	if oneLen <= twoLen {
		num1, num2 = num2, num1
		oneLen, twoLen = twoLen, oneLen
	}
	for i := 0; i < twoLen; i++ {
		curSum := (num1[i]-'0')+(num2[i]-'0')+addBit
		output += strconv.Itoa(int(curSum%10))
		addBit = 0
		if curSum >= 10 {
			addBit = 1
		}
	}
	for i := twoLen; i < oneLen; i++ {
		curSum := num1[i]-'0'+addBit
		output += strconv.Itoa(int(curSum%10))
		addBit = 0
		if curSum >= 10 {
			addBit = 1
		}
	}
	if addBit > 0 {
		output += "1"
	}
	return reverseStr(output)
}
//leetcode submit region end(Prohibit modification and deletion)
