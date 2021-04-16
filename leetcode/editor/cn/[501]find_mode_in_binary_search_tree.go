//给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。 
//
// 假定 BST 有如下定义： 
//
// 
// 结点左子树中所含结点的值小于等于当前结点的值 
// 结点右子树中所含结点的值大于等于当前结点的值 
// 左子树和右子树都是二叉搜索树 
// 
//
// 例如： 
//给定 BST [1,null,2,2], 
//
//    1
//    \
//     2
//    /
//   2
// 
//
// 返回[2]. 
//
// 提示：如果众数超过1个，不需考虑输出顺序 
//
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内） 
// Related Topics 树 
// 👍 294 👎 0


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
func findMode(root *TreeNode) []int {
	tree := make([]int, 0)
	stack := make([]*TreeNode, 0)
	m := make(map[int]int)
	max := 1
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tree = append(tree, node.Val)
		root = node.Right
	}
	for _, v := range tree {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
			if m[v] > max {
				max = m[v]
			}
		}
	}
	for i := 0; i < len(tree); i++ {
		if m[tree[i]] == max {
			res = append(res, tree[i])
			for i < len(tree)-1 && tree[i] == tree[i+1] {
				i++
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
