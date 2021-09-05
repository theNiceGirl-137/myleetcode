//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,null,2,3]
//输出：[1,2,3]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1]
//输出：[1]
// 
//
// 示例 4： 
//
// 
//输入：root = [1,2]
//输出：[1,2]
// 
//
// 示例 5： 
//
// 
//输入：root = [1,null,2]
//输出：[1,2]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 深度优先搜索 二叉树 👍 636 👎 0


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
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	s := make(stack, 0)
	s.push(root)
	for !s.empty() {
		node := s.top()
		s.pop()
		ret = append(ret, node.Val)
		// 先右后左，保证左子树先被遍历
		if node.Right != nil {
			s.push(node.Right)
		}
		if node.Left != nil {
			s.push(node.Left)
		}
	}
	return ret
}

type stack []*TreeNode

func (this *stack)push(x *TreeNode) {
	*this = append(*this, x)
}

func (this *stack)pop() {
	*this = (*this)[:len(*this)-1]
}

func (this *stack)empty() bool {
	return len(*this) == 0
}

func (this *stack)top() *TreeNode {
	return (*this)[len(*this)-1]
}
//leetcode submit region end(Prohibit modification and deletion)
