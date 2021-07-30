//给你一个链表数组，每个链表都已经按升序排列。 
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。 
//
// 
//
// 示例 1： 
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
// 
//
// 示例 2： 
//
// 输入：lists = []
//输出：[]
// 
//
// 示例 3： 
//
// 输入：lists = [[]]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// k == lists.length 
// 0 <= k <= 10^4 
// 0 <= lists[i].length <= 500 
// -10^4 <= lists[i][j] <= 10^4 
// lists[i] 按 升序 排列 
// lists[i].length 的总和不超过 10^4 
// 
// Related Topics 链表 分治 堆（优先队列） 归并排序 
// 👍 1418 👎 0


package leetcode

import "container/heap"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	listNodes := &ListNodes{}
	heap.Init(listNodes)
	for _, v := range lists {
		if v != nil {
			heap.Push(listNodes, v)
		}
	}
	head := &ListNode{}
	idx := head
	for listNodes.Len() > 0 {
		val := heap.Pop(listNodes).(*ListNode)
		idx.Next = val
		if val.Next != nil {
			heap.Push(listNodes, val.Next)
		}
		idx = idx.Next
	}
	return head.Next
}

type ListNodes []*ListNode

func (this *ListNodes) Len() int {
	return len(*this)
}

func (this *ListNodes) Less(i, j int) bool {
	return (*this)[i].Val < (*this)[j].Val
}

func (this *ListNodes) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *ListNodes) Pop() interface{} {
	n := len(*this)
	x := (*this)[n-1]
	*this = (*this)[:n-1]
	return x
}

func (this *ListNodes) Push(x interface{}) {
	*this = append(*this, x.(*ListNode))
}
//leetcode submit region end(Prohibit modification and deletion)
