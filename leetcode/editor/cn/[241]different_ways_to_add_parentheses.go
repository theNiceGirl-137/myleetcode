//给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 *
// 。 
//
// 示例 1: 
//
// 输入: "2-1-1"
//输出: [0, 2]
//解释: 
//((2-1)-1) = 0 
//(2-(1-1)) = 2 
//
// 示例 2: 
//
// 输入: "2*3-4*5"
//输出: [-34, -14, -10, -10, 10]
//解释: 
//(2*(3-(4*5))) = -34 
//((2*3)-(4*5)) = -14 
//((2*(3-4))*5) = -10 
//(2*((3-4)*5)) = -10 
//(((2*3)-4)*5) = 10 
// Related Topics 递归 记忆化搜索 数学 字符串 动态规划 
// 👍 398 👎 0


package leetcode

import (
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func diffWaysToCompute(expression string) []int {
	// 利用分治思想，我们可以将加括号转换为先处理两侧的数学表达式
	ways := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		ch := expression[i]
		if ch == '+' || ch == '-' || ch == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch ch {
						case '+': ways = append(ways, l+r)
						case '-': ways = append(ways, l-r)
						case '*': ways = append(ways, l*r)
					}
				}
			}
		}
	}
	if len(ways) == 0 {
		num, _ := strconv.Atoi(expression)
		ways = append(ways, num)
	}
	return ways
	// 某些被分割的子串可能会出现多次
	//  memoization / 从上到下用分治处理+memoization / 接从下到上用动态规划
}
//leetcode submit region end(Prohibit modification and deletion)
