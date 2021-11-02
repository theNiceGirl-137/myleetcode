//给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的
//根节点的引用。 
//
// 一般来说，删除节点可分为两个步骤： 
//
// 
// 首先找到需要删除的节点； 
// 如果找到了，删除它。 
// 
//
// 说明： 要求算法时间复杂度为 O(h)，h 为树的高度。 
//
// 示例: 
//
// 
//root = [5,3,6,2,4,null,7]
//key = 3
//
//    5
//   / \
//  3   6
// / \   \
//2   4   7
//
//给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
//
//一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
//
//    5
//   / \
//  4   6
// /     \
//2       7
//
//另一个正确答案是 [5,2,6,null,4,null,7]。
//
//    5
//   / \
//  2   6
//   \   \
//    4   7
// 
// Related Topics 树 二叉搜索树 二叉树 👍 519 👎 0


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
func deleteNode450(root *TreeNode, key int) *TreeNode {
	// 空树判断
	if root == nil {
		return nil
	}
	// 找到需要被删除的节点
	if root.Val == key {
		// 没有左子树 让 right 代替 root 的位置
		if root.Left == nil {
			return root.Right
		}
		// 没有右子树 让 left 代替 root 的位置
		if root.Right == nil {
			return root.Left
		}
		// 找后继节点
		// 找比当前节点大的最小节点代替
		next := root.Right
		for next.Left != nil {
			next = next.Left
		}
		root.Right = deleteNode450(root.Right, next.Val)
		root.Val = next.Val
		return root
	}
	if root.Val > key {
		root.Left = deleteNode450(root.Left, key)
	} else {
		root.Right = deleteNode450(root.Right, key)
	}
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
