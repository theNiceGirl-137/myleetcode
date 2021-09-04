//给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。 
//
// 
//
// 示例 1： 
//
// 输入：
//    3
//   / \
//  9  20
//    /  \
//   15   7
//输出：[3, 14.5, 11]
//解释：
//第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。
// 
//
// 
//
// 提示： 
//
// 
// 节点值的范围在32位有符号整数范围内。 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 283 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	if root == nil {
		return ans
	}
	q := make(queue, 0)
	q.push(root)
	for !q.empty() {
		var sum float64
		count := q.size()
		for i := 0; i < count; i++ {
			node := q.front()
			q.pop()
			sum += float64(node.Val)
			if node.Left != nil {
				q.push(node.Left)
			}
			if node.Right != nil {
				q.push(node.Right)
			}
		}
		ans = append(ans, sum/float64(count))
	}
	return ans
}

type queue []*TreeNode

func (this *queue)push(x *TreeNode) {
	*this = append(*this, x)
}

func (this *queue)pop() {
	*this = (*this)[1:]
}

func (this *queue)empty() bool {
	if this.size() == 0 {
		return true
	}
	return false
}

func (this *queue)size() int {
	return len(*this)
}

func (this *queue)front() *TreeNode {
	return (*this)[0]
}
//leetcode submit region end(Prohibit modification and deletion)
