//返回与给定的前序和后序遍历匹配的任何二叉树。 
//
// pre 和 post 遍历中的值是不同的正整数。 
//
// 
//
// 示例： 
//
// 输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
//输出：[1,2,3,4,5,6,7]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= pre.length == post.length <= 30 
// pre[] 和 post[] 都是 1, 2, ..., pre.length 的排列 
// 每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个。 
// 
// Related Topics 树 数组 哈希表 分治 二叉树 👍 189 👎 0


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
func constructFromPrePoast(preorder []int, postorder []int) *TreeNode {
	var tree func(pre []int, post []int) *TreeNode
	tree = func(pre []int, post []int) *TreeNode {
		// 如果左或者右子树为空，直接返回nil
		if len(pre) == 0 {
			return nil
		}
		// 创建节点，赋值
		root := &TreeNode{Val: pre[0]}
		// 考虑到根据 pre[1] 来寻找左右子树
		// 遇到只有一个节点的左右子树，直接返回
		if len(pre) == 1 {
			return root
		}
		// 找左子树的根节点 pre[1]
		// 前序遍历的根节点的下一个节点是左子树的根节点
		mid := 0
		for i, v := range post {
			if v == pre[1] {
				mid = i
				break
			}
		}
		// 递归进行
		root.Left = tree(pre[1:mid+2], post[:mid+1])
		root.Right = tree(pre[mid+2:], post[mid+1:len(post) - 1])
		return root
	}
	return tree(preorder, postorder)
}
//leetcode submit region end(Prohibit modification and deletion)
