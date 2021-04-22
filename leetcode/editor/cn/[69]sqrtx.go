//实现 int sqrt(int x) 函数。 
//
// 计算并返回 x 的平方根，其中 x 是非负整数。 
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。 
//
// 示例 1: 
//
// 输入: 4
//输出: 2
// 
//
// 示例 2: 
//
// 输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842..., 
//     由于返回类型是整数，小数部分将被舍去。
// 
// Related Topics 数学 二分查找 
// 👍 658 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	// 可以将本题理解为求 f(a)=a^2-x=0 的解，而且 f(0)<=0、f(x)>=0
	// 所以本题的解一定在区间 [0,x] 内，故可以用二分法找到解
	l, r := 1, x
	for l <= r {
		mid := l + (r-l)/2
		sqrt := x / mid
		// 判断 sqrt 与 mid 是否相等
		if sqrt == mid {
			return mid
		} else if sqrt < mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r

	// 牛顿迭代法，其公式为 Xn+1=Xn-f(Xn)/f`(Xn)
	// 对于 f(x)=x^2-a=0，迭代公示为 Xn+1=(Xn+a/Xn)/2
	//y := int64(x)
	//x1 := int64(x)
	//for y*y > x1 {
	//	y = (y + x1/y)/2
	//}
	//return int(y)
}
//leetcode submit region end(Prohibit modification and deletion)
