//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。 
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
//输出：3
//解释：和等于 8 的路径有 3 条，如图所示。
// 
//
// 示例 2： 
//
// 
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：3
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [0,1000] 
// -10⁹ <= Node.val <= 10⁹ 
// -1000 <= targetSum <= 1000 
// 
// Related Topics 树 深度优先搜索 二叉树 👍 966 👎 0


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
func pathSum437(root *TreeNode, targetSum int) int {
	//（1）如果选取该节点加入路径，则之后必须继续加入连续节点，或停止加入节点
	//（2）如果不选取该节点加入路径，则对其左右节点进行重新进行考虑
	if root == nil {
		return 0
	}
	return pathSumStartWithRoot(root, targetSum)+pathSum437(root.Left, targetSum)+pathSum437(root.Right, targetSum)
}

func pathSumStartWithRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var count int
	if root.Val == sum {
		count = 1
	}
	count += pathSumStartWithRoot(root.Left, sum-root.Val)
	count += pathSumStartWithRoot(root.Right, sum-root.Val)
	return count
}
//leetcode submit region end(Prohibit modification and deletion)
