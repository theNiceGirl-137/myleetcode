//给定一个二叉树，检查它是否是镜像对称的。 
//
// 
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。 
//
//     1
//   / \
//  2   2
// / \ / \
//3  4 4  3
// 
//
// 
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的: 
//
//     1
//   / \
//  2   2
//   \   \
//   3    3
// 
//
// 
//
// 进阶： 
//
// 你可以运用递归和迭代两种方法解决这个问题吗？ 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1516 👎 0


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
func isSymmetric(root *TreeNode) bool {
	// 判断一个树是否对称等价于判断左右子树是否对称。笔者一般习惯将判断两个子树是否相等
	// 或对称类型的题的解法叫做“四步法”：
	//（1）如果两个子树都为空指针，则它们相等或对称
	//（2）如果两个子树只有一个为空指针，则它们不相等或不对称
	//（3）如果两个子树根节点的值不相等，则它们不相等或不对称
	//（4）根据相等或对称要求，进行递归处理
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
//leetcode submit region end(Prohibit modification and deletion)
