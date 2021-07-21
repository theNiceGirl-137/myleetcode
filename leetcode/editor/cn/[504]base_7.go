//给定一个整数，将其转化为7进制，并以字符串形式输出。 
//
// 示例 1: 
//
// 
//输入: 100
//输出: "202"
// 
//
// 示例 2: 
//
// 
//输入: -7
//输出: "-10"
// 
//
// 注意: 输入范围是 [-1e7, 1e7] 。 
// Related Topics 数学 
// 👍 90 👎 0


package leetcode

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	// 进制转换类型的题，通常是利用除法和取模(mod)来进行计算
	// 同时也要注意一些细节，如负数和零
	// 如果输出是数字类型而非字符串，则也需要考虑是否会超出整数上下界
	if num == 0 {
		return "0"
	}
	isNegative := num < 0
	if isNegative {
		num = -num
	}
	var ans string
	for num > 0 {
		// 从低位向高位拼接
		a := num/7
		b := num%7
		ans = strconv.Itoa(b)+ans
		num = a
	}
	if isNegative {
		ans = "-"+ans
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
