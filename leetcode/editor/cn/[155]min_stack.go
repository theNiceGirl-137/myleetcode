//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。 
//
// 
// push(x) —— 将元素 x 推入栈中。 
// pop() —— 删除栈顶的元素。 
// top() —— 获取栈顶元素。 
// getMin() —— 检索栈中的最小元素。 
// 
//
// 
//
// 示例: 
//
// 输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
// 
//
// 
//
// 提示： 
//
// 
// pop、top 和 getMin 操作总是在 非空栈 上调用。 
// 
// Related Topics 栈 设计 
// 👍 966 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	s, minS []int
}


/** initialize your data structure here. */
func Constructor155() MinStack {
	s := make([]int, 0)
	minS := make([]int, 0)
	return MinStack{s: s, minS: minS}
}


func (this *MinStack) Push(val int)  {
	this.s = push(this.s, val)
	if len(this.minS) == 0 || top(this.minS) >= val {
		this.minS = push(this.minS, val)
	}
}


func (this *MinStack) Pop()  {
	if len(this.minS) != 0 && top(this.minS) == top(this.s) {
		this.minS = pop(this.minS)
	}
	this.s = pop(this.s)
}


func (this *MinStack) Top() int {
	return top(this.s)
}


func (this *MinStack) GetMin() int {
	return top(this.minS)
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
