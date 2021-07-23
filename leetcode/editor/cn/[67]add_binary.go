//给你两个二进制字符串，返回它们的和（用二进制表示）。 
//
// 输入为 非空 字符串且只包含数字 1 和 0。 
//
// 
//
// 示例 1: 
//
// 输入: a = "11", b = "1"
//输出: "100" 
//
// 示例 2: 
//
// 输入: a = "1010", b = "1011"
//输出: "10101" 
//
// 
//
// 提示： 
//
// 
// 每个字符串仅由字符 '0' 或 '1' 组成。 
// 1 <= a.length, b.length <= 10^4 
// 字符串如果不是 "0" ，就都不含前导零。 
// 
// Related Topics 位运算 数学 字符串 模拟 
// 👍 647 👎 0


package leetcode

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	var output string
	var addBit uint8
	a, b = reverseStr67(a), reverseStr67(b)
	aLen, bLen := len(a), len(b)
	if len(a) <= len(b) {
		a, b = b, a
		aLen, bLen = bLen, aLen
	}
	for i := 0; i < bLen; i++ {
		curSum := (a[i]-'0')+(b[i]-'0')+addBit
		output += strconv.Itoa(int(curSum%2))
		addBit = 0
		if curSum >= 2 {
			addBit = 1
		}
	}
	for i := bLen; i < aLen; i++ {
		curSum := (a[i]-'0')+addBit
		output += strconv.Itoa(int(curSum%2))
		addBit = 0
		if curSum >= 2 {
			addBit = 1
		}
	}
	if addBit > 0 {
		output += "1"
	}
	return reverseStr67(output)
}

func reverseStr67(str string) string {
	r := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
//leetcode submit region end(Prohibit modification and deletion)
