//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。 
//
// 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bs
//t/ 相同 
//
// 
//
// 
// 
// 示例 1： 
//
// 
//输入：root = [4,2,6,1,3]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：root = [1,0,48,null,null,12,49]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [2, 100] 内 
// 0 <= Node.val <= 105 
// 
// 
// 
// Related Topics 树 深度优先搜索 递归 
// 👍 142 👎 0


package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func minDiffInBST(root *TreeNode) int {
	var pre []int
	stack := make([]*TreeNode, 0)
	min := math.MaxInt32
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pre = append(pre, node.Val)
		if len(pre) > 1 {
			sub := node.Val - pre[len(pre)-2]
			if min > sub {
				min = sub
			}
		}
		root = node.Right
	}
	if len(pre) < 2 {
		return pre[0]
	}
	return min
}
//leetcode submit region end(Prohibit modification and deletion)
