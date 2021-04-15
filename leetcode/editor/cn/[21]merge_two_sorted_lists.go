//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 
//
// 示例 1： 
//
// 
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
// 
//
// 示例 2： 
//
// 
//输入：l1 = [], l2 = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [], l2 = [0]
//输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 两个链表的节点数目范围是 [0, 50] 
// -100 <= Node.val <= 100 
// l1 和 l2 均按 非递减顺序 排列 
// 
// Related Topics 递归 链表 
// 👍 1668 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 遍历两个链表依次拼接，增加了哨兵节点
	//pre := &ListNode{}
	//newHead:= pre
	//for l1 != nil && l2 != nil {
	//	if l1.Val < l2.Val {
	//		pre.Next = l1
	//		l1 = l1.Next
	//	} else {
	//		pre.Next = l2
	//		l2 = l2.Next
	//	}
	//	pre = pre.Next
	//}
	//if l1 != nil {
	//	pre.Next = l1
	//}
	//if l2 != nil {
	//	pre.Next = l2
	//}
	//return newHead.Next

	// 递归解法
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newHead := &ListNode{}
	if l1.Val < l2.Val {
		newHead = l1
		newHead.Next = mergeTwoLists(l1.Next, l2)
	} else {
		newHead = l2
		newHead.Next = mergeTwoLists(l2.Next, l1)
	}
	return newHead
}
//leetcode submit region end(Prohibit modification and deletion)
