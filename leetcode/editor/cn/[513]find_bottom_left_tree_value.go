//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。 
//
// 假设二叉树中至少有一个节点。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: root = [2,1,3]
//输出: 1
// 
//
// 示例 2: 
//
// 
//
// 
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [1,10⁴] 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 198 👎 0


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
var maxDeep, res int
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	findBottomLeftValueHelper(root, maxDeep)
	return res
}

func findBottomLeftValueHelper(root *TreeNode, deep int) {
	if root.Left == nil && root.Right == nil {
		if deep > maxDeep {
			res = root.Val
			maxDeep = deep
		}
	}
	if root.Left != nil {
		deep++
		findBottomLeftValueHelper(root.Left, deep)
		deep--
	}
	if root.Right != nil {
		deep++
		findBottomLeftValueHelper(root.Right, deep)
		deep--
	}
}
//leetcode submit region end(Prohibit modification and deletion)
