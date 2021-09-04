//给出二叉树的根节点 root，树上每个节点都有一个不同的值。 
//
// 如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。 
//
// 返回森林中的每棵树。你可以按任意顺序组织答案。 
//
// 
//
// 示例： 
//
// 
//
// 输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
//输出：[[1,2,null,4],[6],[7]]
// 
//
// 
//
// 提示： 
//
// 
// 树中的节点数最大为 1000。 
// 每个节点都有一个介于 1 到 1000 之间的值，且各不相同。 
// to_delete.length <= 1000 
// to_delete 包含一些从 1 到 1000、各不相同的值。 
// 
// Related Topics 树 深度优先搜索 二叉树 👍 135 👎 0


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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    forest := make([]*TreeNode, 0)
    dict := make(map[int]bool)

    for i := 0; i < len(to_delete); i++ {
    	dict[to_delete[i]] = true
	}

	root = helper1110(root, dict, &forest)
	if root != nil {
		forest = append([]*TreeNode{root}, forest...)
	}
	return forest
}

func helper1110(root *TreeNode, dict map[int]bool, forest *[]*TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = helper1110(root.Left, dict, forest)
	root.Right = helper1110(root.Right, dict, forest)
	if dict[root.Val] == true {
		if root.Right != nil {
			*forest = append([]*TreeNode{root.Right}, *forest...)
		}
		if root.Left != nil {
			*forest = append([]*TreeNode{root.Left}, *forest...)
		}
		root = nil
	}
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
