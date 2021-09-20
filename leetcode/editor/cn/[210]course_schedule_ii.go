//现在你总共有 n 门课需要选，记为 0 到 n-1。 
//
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1] 
//
// 给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。 
//
// 可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。 
//
// 示例 1: 
//
// 输入: 2, [[1,0]] 
//输出: [0,1]
//解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。 
//
// 示例 2: 
//
// 输入: 4, [[1,0],[2,0],[3,1],[3,2]]
//输出: [0,1,2,3] or [0,2,1,3]
//解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
//     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
// 
//
// 说明: 
//
// 
// 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。 
// 你可以假定输入的先决条件中没有重复的边。 
// 
//
// 提示: 
//
// 
// 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。 
// 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。 
// 
// 拓扑排序也可以通过 BFS 完成。 
// 
// 
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 475 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 我们可以先建立一个邻接矩阵表示图，方便进行直接查找
	// 这里注意我们将所有的边反向，使得如果课程 i 指向课程 j，那么课程 i 需要在课程 j 前面先修完，这样更符合我们的直观理解
	// 拓扑排序也可以被看成是广度优先搜索的一种情况：我们先遍历一遍所有节点，把入度为 0 的节点（即没有前置课程要求）放在队列中
	// 在每次从队列中获得节点时，我们将该节点放在目前排序的末尾，并且把它指向的课程的入度各减 1
	// 如果在这个过程中有课程的所有前置必修课都已修完（即入度为 0），我们把这个节点加入队列中
	// 当队列的节点都被处理完时，说明所有的节点都已排好序，或因图中存在循环而无法上完所有课程
	graph := make([][]int, numCourses)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}
	res := make([]int, 0)
	indegree := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		indegree[prerequisite[0]]++
	}
	q := make(queue, 0)
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			q.push(i)
		}
	}
	for !q.empty() {
		u := q.front()
		res = append(res, u)
		q.pop()
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				q.push(v)
			}
		}
	}
	for i := 0; i < len(indegree); i++ {
		if indegree[i] != 0 {
			return []int{}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
