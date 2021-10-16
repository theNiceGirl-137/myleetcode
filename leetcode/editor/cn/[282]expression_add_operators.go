//给定一个仅包含数字 0-9 的字符串 num 和一个目标值整数 target ，在 num 的数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回所
//有能够得到目标值的表达式。 
//
// 
//
// 示例 1: 
//
// 
//输入: num = "123", target = 6
//输出: ["1+2+3", "1*2*3"] 
// 
//
// 示例 2: 
//
// 
//输入: num = "232", target = 8
//输出: ["2*3+2", "2+3*2"] 
//
// 示例 3: 
//
// 
//输入: num = "105", target = 5
//输出: ["1*0+5","10-5"] 
//
// 示例 4: 
//
// 
//输入: num = "00", target = 0
//输出: ["0+0", "0-0", "0*0"]
// 
//
// 示例 5: 
//
// 
//输入: num = "3456237490", target = 9191
//输出: [] 
//
// 
//
// 提示： 
//
// 
// 1 <= num.length <= 10 
// num 仅含数字 
// -2³¹ <= target <= 2³¹ - 1 
// 
// Related Topics 数学 字符串 回溯 👍 284 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func addOperators(num string, target int) []string {
	n := len(num)
	ans := make([]string, 0)
	backTracking282(num, &ans, make([]byte, 0, n*2-1), 0, 0, 0, n, target)
	return ans
}

func backTracking282(num string, ans *[]string, expr []byte, i, res, mul, n, target int) {
	if i == n {
		if res == target {
			*ans = append(*ans, string(expr))
		}
		return
	}
	signIndex := len(expr)
	if i > 0 {
		// 占位，准备填充符号
		expr = append(expr, 0)
	}
	// 枚举截取的数字长度（取多少位）
	// 数字不能有前导 0
	for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
		val = val*10 + int(num[j]-'0')
		expr = append(expr, num[j])
		if i == 0 {
			backTracking282(num, ans, expr, j+1, val, val, n, target)
		} else {
			expr[signIndex] = '+'
			backTracking282(num, ans, expr, j+1, res+val, val, n, target)
			expr[signIndex] = '-'
			backTracking282(num, ans, expr, j+1, res-val, -val, n, target)
			expr[signIndex] = '*'
			backTracking282(num, ans, expr, j+1, res-mul+mul*val, mul*val, n, target)
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
