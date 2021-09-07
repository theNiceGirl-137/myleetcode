//计算给定二叉树的所有左叶子之和。 
//
// 示例： 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24 
//
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 341 👎 0


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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	sumOfLeftLeavesHelper(root, &res)
	return res
}

func sumOfLeftLeavesHelper(root *TreeNode, res *int) {
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
	}
	if root.Left != nil {
		sumOfLeftLeavesHelper(root.Left, res)
	}
	if root.Right != nil {
		sumOfLeftLeavesHelper(root.Right, res)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
