//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。 
//
// 整数除法仅保留整数部分。 
//
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：s = "3+2*2"
//输出：7
// 
//
// 示例 2： 
//
// 
//输入：s = " 3/2 "
//输出：1
// 
//
// 示例 3： 
//
// 
//输入：s = " 3+5 / 2 "
//输出：5
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 3 * 10⁵ 
// s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开 
// s 表示一个 有效表达式 
// 表达式中的所有整数都是非负整数，且在范围 [0, 2³¹ - 1] 内 
// 题目数据保证答案是一个 32-bit 整数 
// 
// 
// 
// Related Topics 栈 数学 字符串 👍 446 👎 0


package leetcode

import "unicode"

//leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	return parseExpr(s, 0)
}

// parseExpr 递归 parse 从位置 i 开始的剩余字符串
func parseExpr(s string, i int) int {
	var left, right int
	op := '+'
	for i < len(s) {
		if s[i] != ' ' {
			n := parseNum(s, &i)
			switch op {
			case '+':
				left += right
				right = n
			case '-':
				left += right
				right = -n
			case '*':
				right *= n
			case '/':
				right /= n
			}
			if i < len(s) {
				op = int32(s[i])
			}
		}
		i++
	}
	return left+right
}

// parseNum parse 从位置 i 开始的一个数字
func parseNum(s string, i *int) int {
	var n int
	for *i < len(s) && unicode.IsDigit(rune(s[*i])) {
		n = 10*n+int(s[*i]-'0')
		*i++
	}
	return n
}
//leetcode submit region end(Prohibit modification and deletion)
