//给定一个二叉树，返回所有从根节点到叶子节点的路径。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例: 
//
// 输入:
//
//   1
// /   \
//2     3
// \
//  5
//
//输出: ["1->2->5", "1->3"]
//
//解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3 
// Related Topics 树 深度优先搜索 
// 👍 495 👎 0


package leetcode

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var path []string
	if root == nil {
		return ans
	}
	dfs257(root, &ans, path)
	return ans
}

func dfs257(root *TreeNode, ans *[]string, path []string) {
	if root == nil {
		return
	}
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		temp := make([]string, len(path))
		copy(temp, path)
		*ans = append(*ans, strings.Join(temp, "->"))
		return
	}
	if root.Left != nil {
		dfs257(root.Left, ans, path)
	}
	if root.Right != nil {
		dfs257(root.Right, ans, path)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
