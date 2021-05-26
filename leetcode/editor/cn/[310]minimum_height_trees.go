//树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。 
//
// 给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），其中
// edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。 
//
// 可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度
//树 。 
//
// 请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。 
//树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。
//
// 
//
// 示例 1： 
//
// 
//输入：n = 4, edges = [[1,0],[1,2],[1,3]]
//输出：[1]
//解释：如图所示，当根是标签为 1 的节点时，树的高度是 1 ，这是唯一的最小高度树。 
//
// 示例 2： 
//
// 
//输入：n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
//输出：[3,4]
// 
//
// 示例 3： 
//
// 
//输入：n = 1, edges = []
//输出：[0]
// 
//
// 示例 4： 
//
// 
//输入：n = 2, edges = [[0,1]]
//输出：[0,1]
// 
//
// 
//
// 
// 
//
// 提示： 
//
// 
// 1 <= n <= 2 * 104 
// edges.length == n - 1 
// 0 <= ai, bi < n 
// ai != bi 
// 所有 (ai, bi) 互不相同 
// 给定的输入 保证 是一棵树，并且 不会有重复的边 
// 
// Related Topics 广度优先搜索 图 
// 👍 329 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func findMinHeightTrees(n int, edges [][]int) []int {
	var res []int
	if n == 1 {
		return append(res, 0)
	}
	// 统计节点出度
	degree := make([]int, n)
	// 统计节点之间的关系
	edgeMap := make([][]int, n)
	for _, edge := range edges {
		degree[edge[0]]++
		degree[edge[1]]++
		edgeMap[edge[0]] = append(edgeMap[edge[0]], edge[1])
		edgeMap[edge[1]] = append(edgeMap[edge[1]], edge[0])
	}
	queue := make([]int, 0)
	// 把所有出度为 1 的节点入队，即叶子节点
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		// 每次循环都新建一个结果集，保证最后一定是最小高度树的节点集合
		res = make([]int, 0)
		// 记录每一层的节点数
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			res = append(res, cur)
			neighbors := edgeMap[cur]
			for _, neighbor := range neighbors {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)