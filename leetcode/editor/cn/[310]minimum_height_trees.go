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

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
var direction310 = []int{-1, 0, 1, 0, -1}

func findMinHeightTrees(n int, edges [][]int) []int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})
	return nil
}
//leetcode submit region end(Prohibit modification and deletion)
//import java.util.ArrayList;
//import java.util.LinkedList;
//import java.util.List;
//import java.util.Queue;
//
//public class FindMinHeightTrees {
//public List<Integer> findMinHeightTrees(int n, int[][] edges) {
//List<Integer> res = new ArrayList<>();
///*如果只有一个节点，那么他就是最小高度树*/
//if (n == 1) {
//res.add(0);
//return res;
//}
///*建立各个节点的出度表*/
//int[] degree = new int[n];
///*建立图关系，在每个节点的list中存储相连节点*/
//List<List<Integer>> map = new ArrayList<>();
//for (int i = 0; i < n; i++) {
//map.add(new ArrayList<>());
//}
//for (int[] edge : edges) {
//degree[edge[0]]++;
//degree[edge[1]]++;/*出度++*/
//map.get(edge[0]).add(edge[1]);/*添加相邻节点*/
//map.get(edge[1]).add(edge[0]);
//}
///*建立队列*/
//Queue<Integer> queue = new LinkedList<>();
///*把所有出度为1的节点，也就是叶子节点入队*/
//for (int i = 0; i < n; i++) {
//if (degree[i] == 1) queue.offer(i);
//}
///*循环条件当然是经典的不空判断*/
//while (!queue.isEmpty()) {
//res = new ArrayList<>();/*这个地方注意，我们每层循环都要new一个新的结果集合，
//  这样最后保存的就是最终的最小高度树了*/
//int size = queue.size();/*这是每一层的节点的数量*/
//for (int i = 0; i < size; i++) {
//int cur = queue.poll();
//res.add(cur);/*把当前节点加入结果集，不要有疑问，为什么当前只是叶子节点为什么要加入结果集呢?
//  因为我们每次循环都会新建一个list，所以最后保存的就是最后一个状态下的叶子节点，
//  这也是很多题解里面所说的剪掉叶子节点的部分，你可以想象一下图，每层遍历完，
//  都会把该层（也就是叶子节点层）这一层从队列中移除掉，
//  不就相当于把原来的图给剪掉一圈叶子节点，形成一个缩小的新的图吗*/
//List<Integer> neighbors = map.get(cur);
///*这里就是经典的bfs了，把当前节点的相邻接点都拿出来，
// * 把它们的出度都减1，因为当前节点已经不存在了，所以，
// * 它的相邻节点们就有可能变成叶子节点*/
//for (int neighbor : neighbors) {
//degree[neighbor]--;
//if (degree[neighbor] == 1) {
///*如果是叶子节点我们就入队*/
//queue.offer(neighbor);
//}
//}
//}
//}
//return res;/*返回最后一次保存的list*/
//}
//
//}
