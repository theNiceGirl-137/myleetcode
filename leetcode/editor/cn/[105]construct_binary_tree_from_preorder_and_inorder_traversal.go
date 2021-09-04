//给定一棵树的前序遍历 preorder 与中序遍历 inorder。请构造二叉树并返回其根节点。 
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
// 提示: 
//
// 
// 1 <= preorder.length <= 3000 
// inorder.length == preorder.length 
// -3000 <= preorder[i], inorder[i] <= 3000 
// preorder 和 inorder 均无重复元素 
// inorder 均出现在 preorder 
// preorder 保证为二叉树的前序遍历序列 
// inorder 保证为二叉树的中序遍历序列 
// 
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1191 👎 0


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
	if len(preorder) == 0 {
		return nil
	}
	hash := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		hash[inorder[i]] = i
	}
	return buildTreeHelper(hash, preorder, 0, len(preorder)-1, 0)
}

func buildTreeHelper(hash map[int]int, preorder []int, s0, e0, s1 int) *TreeNode {
	if s0 > e0 {
		return nil
	}
	mid := preorder[s1]
	index := hash[mid]
	leftLen := index-s0-1
	node := &TreeNode{Val: mid}
	node.Left = buildTreeHelper(hash, preorder, s0, index-1, s1+1)
	node.Right = buildTreeHelper(hash, preorder, index+1, e0, s1+2+leftLen)
	return node
}
//leetcode submit region end(Prohibit modification and deletion)
