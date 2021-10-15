//输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。 
//
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。 
//
// 
//
// 示例 1: 
//
// 
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
// 
//
// 示例 2: 
//
// 
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
// 
//
// 
//
// 限制： 
//
// 0 <= 节点个数 <= 5000 
//
// 
//
// 注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-
//preorder-and-inorder-traversal/ 
// Related Topics 树 数组 哈希表 分治 二叉树 👍 583 👎 0


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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	inorderIndex := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inorderIndex[inorder[i]] = i
	}
	return buildTreeHelper(inorderIndex, preorder, 0, 0, len(preorder)-1)
}

func buildTreeHelper(inorderIndex map[int]int, preorder []int, preRootIndex, leftSide, rightSide int) *TreeNode {
	if leftSide > rightSide {
		return nil
	}
	rootValue := preorder[preRootIndex]
	inRootIndex := inorderIndex[rootValue]
	leftTreeLen := inRootIndex-leftSide
	return &TreeNode{
		Val: rootValue,
		Left: buildTreeHelper(inorderIndex, preorder, preRootIndex+1, leftSide, inRootIndex-1),
		Right: buildTreeHelper(inorderIndex, preorder, preRootIndex+1+leftTreeLen, inRootIndex+1, rightSide),
	}
}
//leetcode submit region end(Prohibit modification and deletion)
