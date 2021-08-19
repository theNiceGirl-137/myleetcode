//给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。 
//
// 进阶: 
//如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？ 
//
// 示例: 
//
// 
//// 初始化一个单链表 [1,2,3].
//ListNode head = new ListNode(1);
//head.next = new ListNode(2);
//head.next.next = new ListNode(3);
//Solution solution = new Solution(head);
//
//// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
//solution.getRandom();
// 
// Related Topics 水塘抽样 链表 数学 随机化 
// 👍 141 👎 0


package leetcode

import "math/rand"

// Solution leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	Head *ListNode
}

// Constructor
// @Param head The linked list's head.
func Constructor382(head *ListNode) Solution {
	return Solution{
		Head: head,
	}
}

// GetRandom  Returns a random node's value.
func (this *Solution) GetRandom() int {
	// 不同于数组，在未遍历完链表前，我们无法知道链表的总长度
	// 这里我们就可以使用水库采样：遍历一次链表，在遍历到第 m 个节点时，有 1/m 的概率选择这个节点覆盖掉之前的节点选择
	// 对于长度为 n 的链表的第 m 个节点，最后被采样的充要条件是它被选择，且之后的节点都没有被选择。这种情况发生的概率为
	// 1/m × m/m+1 × m+1/m+2 × · · · × n−1/n = 1/n
	// 因此每个点都有均等的概率被选择
	ans := this.Head.Val
	node := this.Head.Next
	i := 2
	for node != nil {
		if rand.Int()%i == 0 {
			ans = node.Val
		}
		i++
		node = node.Next
	}
	return ans
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
