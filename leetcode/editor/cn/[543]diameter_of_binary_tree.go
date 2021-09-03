//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。 
//
// 
//
// 示例 : 
//给定二叉树 
//
//           1
//         / \
//        2   3
//       / \     
//      4   5    
// 
//
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。 
//
// 
//
// 注意：两结点之间的路径长度是以它们之间边的数目表示。 
// Related Topics 树 深度优先搜索 二叉树 👍 782 👎 0


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
func diameterOfBinaryTree(root *TreeNode) int {
	// 同样的，我们可以利用递归来处理树。解题时要注意，在我们处理某个子树时，我们更新的
	// 最长直径值和递归返回的值是不同的。这是因为待更新的最长直径值是经过该子树根节点的最长
	// 直径（即两侧长度）；而函数返回值是以该子树根节点为端点的最长直径值（即一侧长度），使用
	// 这样的返回值才可以通过递归更新父节点的最长直径值
	var diameter int
	helper543(root, &diameter)
	return diameter
}

func helper543(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left, right := helper543(root.Left, diameter), helper543(root.Right, diameter)
	*diameter = max(left+right, *diameter)
	return max(left, right)+1
}
//leetcode submit region end(Prohibit modification and deletion)
