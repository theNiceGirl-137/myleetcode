//根据一棵树的中序遍历与后序遍历构造二叉树。 
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
// Related Topics 树 深度优先搜索 数组 
// 👍 485 👎 0


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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 后序遍历的结果的最后一位是二叉树的根
	// 跟树有关的题目可以考虑使用递归
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal, rootLoc := postorder[len(postorder)-1], 0
	for i := range inorder {
		if rootVal == inorder[i] {
			rootLoc = i
			break
		}
	}
	// 中序遍历根节点左边的部分为左子树，右边的部分为右子树 （中序 左 中 右）
	// 后序遍历根据中序遍历确定的根节点位置判断左子树节点个数（后序 左 右 中）
	return &TreeNode {
		Val: rootVal,
		Left: buildTree(inorder[:rootLoc], postorder[:rootLoc]),
		Right: buildTree(inorder[rootLoc+1:], postorder[rootLoc:len(postorder)-1]),
	}
}
//leetcode submit region end(Prohibit modification and deletion)
