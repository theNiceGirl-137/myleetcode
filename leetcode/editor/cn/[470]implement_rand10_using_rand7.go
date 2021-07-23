//已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。 
//
// 不要使用系统的 Math.random() 方法。 
//
// 
// 
//
// 
//
// 示例 1: 
//
// 
//输入: 1
//输出: [7]
// 
//
// 示例 2: 
//
// 
//输入: 2
//输出: [8,4]
// 
//
// 示例 3: 
//
// 
//输入: 3
//输出: [8,1,10]
// 
//
// 
//
// 提示: 
//
// 
// rand7 已定义。 
// 传入参数: n 表示 rand10 的调用次数。 
// 
//
// 
//
// 进阶: 
//
// 
// rand7()调用次数的 期望值 是多少 ? 
// 你能否尽量少调用 rand7() ? 
// 
// Related Topics 数学 拒绝采样 概率与统计 随机化 
// 👍 209 👎 0


package leetcode

import "math/rand"

//leetcode submit region begin(Prohibit modification and deletion)
func rand10() int {
	// X、Y为正整数
	// 规律1：(randX() - 1) * Y + randY() = randX*Y()
	// 规律2： randX*Y() % 10 + 1 = randY()
	var a, b int
	for {
		a, b = rand7(), rand7()
		rand49 := (a-1)*7+b
		if rand49 <= 40 {
			return rand49%10+1
		}
		a, b = rand49-40, rand7() // a --> rand9
		rand63 := (a-1)*7+b
		if rand63 <= 60 {
			return rand63%10+1
		}
		a, b = rand63-60, rand7() // a --> rand3
		rand21 := (a-1)*7+b
		if rand21 <= 20 {
			return rand21%10+1
		}
	}
}

func rand7() int {
	return rand.Intn(7)
}
//leetcode submit region end(Prohibit modification and deletion)
