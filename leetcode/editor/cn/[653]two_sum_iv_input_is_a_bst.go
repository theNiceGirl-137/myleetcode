//给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。 
//
// 
//
// 示例 1： 
//
// 
//输入: root = [5,3,6,2,4,null,7], k = 9
//输出: true
// 
//
// 示例 2： 
//
// 
//输入: root = [5,3,6,2,4,null,7], k = 28
//输出: false
// 
//
// 示例 3： 
//
// 
//输入: root = [2,1,3], k = 4
//输出: true
// 
//
// 示例 4： 
//
// 
//输入: root = [2,1,3], k = 1
//输出: false
// 
//
// 示例 5： 
//
// 
//输入: root = [2,1,3], k = 3
//输出: true
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [1, 10⁴]. 
// -10⁴ <= Node.val <= 10⁴ 
// root 为二叉搜索树 
// -10⁵ <= k <= 10⁵ 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 哈希表 双指针 二叉树 👍 276 👎 0


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
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return false
	}
	nums := make(map[int]bool)
	var getNums func(t *TreeNode)
	getNums = func(t *TreeNode) {
		if t == nil {
			return
		}
		nums[t.Val] = true
		getNums(t.Left)
		getNums(t.Right)
	}
	getNums(root)
	for num, _ := range nums {
		if nums[k-num] == true && k-num != num {
			return true
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
