//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "()"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "()[]{}"
//输出：true
// 
//
// 示例 3： 
//
// 
//输入：s = "(]"
//输出：false
// 
//
// 示例 4： 
//
// 
//输入：s = "([)]"
//输出：false
// 
//
// 示例 5： 
//
// 
//输入：s = "{[]}"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 104 
// s 仅由括号 '()[]{}' 组成 
// 
// Related Topics 栈 字符串 
// 👍 2522 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	// 典型需要用栈来解决的问题
	temp := []byte(s)
	stack := make([]byte, 0)
	for i := 0; i < len(temp); i++ {
		if temp[i] == '(' || temp[i] == '[' || temp[i] == '{' {
			stack = append(stack, temp[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			ch := sTop(stack)
			if ch == '(' && temp[i] == ')' || ch == '[' && temp[i] == ']' || ch == '{' && temp[i] == '}' {
				stack = sPop(stack)
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func sTop(arr []byte) byte {
	return arr[len(arr)-1]
}

func sPop(arr []byte) []byte {
	return arr[:len(arr)-1]
}

func sPush(arr []byte, x byte) []byte {
	return append(arr, x)
}
//leetcode submit region end(Prohibit modification and deletion)
