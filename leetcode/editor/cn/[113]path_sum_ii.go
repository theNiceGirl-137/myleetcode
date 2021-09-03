//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。 
//
// 叶子节点 是指没有子节点的节点。 
//
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
// 
//
// 示例 2： 
//
// 
//输入：root = [1,2,3], targetSum = 5
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1,2], targetSum = 0
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点总数在范围 [0, 5000] 内 
// -1000 <= Node.val <= 1000 
// -1000 <= targetSum <= 1000 
// 
// 
// 
// Related Topics 树 深度优先搜索 
// 👍 469 👎 0


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
func pathSum113(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	dfs113(root, targetSum, path, &result)
	return result
}

func dfs113(root *TreeNode, targetSum int, path []int, result *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			// 数组是一种值
			// 每次 append 之后不一定返回一个新切片地址，所以需要 copy 来保证数据准确
			// 将 path 添加进结果集时先复制再加入，直接操作原切片数据会被后面的遍历修改
			newPath := make([]int, len(path)+1)
			newPath[len(path)] = root.Val
			copy(newPath, path)
			*result = append(*result, newPath)
		}
		return
	}
	dfs113(root.Left, targetSum-root.Val, append(path, root.Val), result)
	dfs113(root.Right, targetSum-root.Val, append(path, root.Val), result)
}
//leetcode submit region end(Prohibit modification and deletion)
