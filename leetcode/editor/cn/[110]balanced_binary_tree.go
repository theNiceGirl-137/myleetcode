//给定一个二叉树，判断它是否是高度平衡的二叉树。 
//
// 本题中，一棵高度平衡二叉树定义为： 
//
// 
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：true
// 
//
// 
//
// 提示： 
//
// 
// 树中的节点数在范围 [0, 5000] 内 
// -10⁴ <= Node.val <= 10⁴ 
// 
// Related Topics 树 深度优先搜索 二叉树 👍 756 👎 0


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
func isBalanced(root *TreeNode) bool {
	// 解法类似于求树的最大深度，但有两个不同的地方：一是我们需要先处理子树的深度再进行
	// 比较，二是如果我们在处理子树时发现其已经不平衡了，则可以返回一个-1，使得所有其长辈节
	// 点可以避免多余的判断
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := helper(root.Left), helper(root.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right)+1
}
//leetcode submit region end(Prohibit modification and deletion)
