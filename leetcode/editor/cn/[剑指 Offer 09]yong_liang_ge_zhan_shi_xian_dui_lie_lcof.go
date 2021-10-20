//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 ) 
//
// 
//
// 示例 1： 
//
// 输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
// 
//
// 示例 2： 
//
// 输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
// 
//
// 提示： 
//
// 
// 1 <= values <= 10000 
// 最多会对 appendTail、deleteHead 进行 10000 次调用 
// 
// Related Topics 栈 设计 队列 👍 332 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
	in, out stack
}


func ConstructorOffer08() CQueue {
	return CQueue{
		in: make(stack, 0),
		out: make(stack, 0),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.in.push(value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		n := len(this.in)
		for i := 0; i < n; i++ {
			x := this.in.top()
			this.out.push(x)
			this.in.pop()
		}
	}
	if len(this.out) == 0 {
		return -1
	}
	v := this.out.top()
	this.out.pop()
	return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)
