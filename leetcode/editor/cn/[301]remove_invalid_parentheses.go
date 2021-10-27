//给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。 
//
// 返回所有可能的结果。答案可以按 任意顺序 返回。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "()())()"
//输出：["(())()","()()()"]
// 
//
// 示例 2： 
//
// 
//输入：s = "(a)())()"
//输出：["(a())()","(a)()()"]
// 
//
// 示例 3： 
//
// 
//输入：s = ")("
//输出：[""]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 25 
// s 由小写英文字母以及括号 '(' 和 ')' 组成 
// s 中至多含 20 个括号 
// 
// Related Topics 广度优先搜索 字符串 回溯 👍 538 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func removeInvalidParentheses(s string) []string {
	// 广度优先搜索
	ans := make([]string, 0)
	curSet := map[string]bool{s: true}
	// 每一轮删除字符中的一个括号，直到出现合法的字符串为止
	// 循环的次数就是需要删除的括号的个数
	for {
		for str := range curSet {
			if isValid301(str) {
				ans = append(ans, str)
			}
		}
		if len(ans) > 0 {
			return ans
		}
		nextSet := map[string]bool{}
		for str := range curSet {
			for i, ch := range str {
				if ch == '(' || ch == ')' {
					nextSet[str[:i]+str[i+1:]] = true
				}
			}
		}
		curSet = nextSet
	}
}

func isValid301(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		}
		if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
//leetcode submit region end(Prohibit modification and deletion)
