//编写一个算法来判断一个数 n 是不是快乐数。 
//
// 「快乐数」定义为： 
//
// 
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。 
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。 
// 如果 可以变为 1，那么这个数就是快乐数。 
// 
//
// 如果 n 是快乐数就返回 true ；不是，则返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//输入：19
//输出：true
//解释：
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1
// 
//
// 示例 2： 
//
// 
//输入：n = 2
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 231 - 1 
// 
// Related Topics 哈希表 数学 双指针 
// 👍 638 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	// 对于平方数链，有两种情况：
	// 要不经过多次计算到达一，因为一的平方还等于一，所以链条到此时进入无限循环，我们可以认为它终止于一，也可以看作是一个不动点
	// 另一种情况是，无论怎么计算都不会到达一，但是会进入另外一个循环：
	// 4→16→37→58→89→145→42→20→4
	// 快慢指针
	slow, fast := n, getSum(n)
	for fast != 1 && slow != fast {
		// 慢指针走一步
		slow = getSum(slow)
		// 快指针走两步
		fast = getSum(getSum(fast))
	}
	return fast == 1
}

func getSum(n int) (sum int) {
	for n > 0 {
		sum += (n%10)*(n%10)
		n /= 10
	}
	return sum
}
//leetcode submit region end(Prohibit modification and deletion)
