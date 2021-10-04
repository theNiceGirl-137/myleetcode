//给你一个无向图（原始图），图中有 n 个节点，编号从 0 到 n - 1 。你决定将图中的每条边细分为一条节点链，每条边之间的新节点数各不相同。 
//
// 图用由边组成的二维数组 edges 表示，其中 edges[i] = [ui, vi, cnti] 表示原始图中节点 ui 和 vi 之间存在一条边，
//cnti 是将边细分后的新节点总数。注意，cnti == 0 表示边不可细分。 
//
// 要细分边 [ui, vi] ，需要将其替换为 (cnti + 1) 条新边，和 cnti 个新节点。新节点为 x1, x2, ..., xcnti ，新边
//为 [ui, x1], [x1, x2], [x2, x3], ..., [xcnti+1, xcnti], [xcnti, vi] 。 
//
// 现在得到一个新的 细分图 ，请你计算从节点 0 出发，可以到达多少个节点？节点 是否可以到达的判断条件 为：如果节点间距离是 maxMoves 或更少，则
//视为可以到达；否则，不可到达。 
//
// 给你原始图和 maxMoves ，返回新的细分图中从节点 0 出发 可到达的节点数 。 
//
// 
//
// 示例 1： 
//
// 
//输入：edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
//输出：13
//解释：边的细分情况如上图所示。
//可以到达的节点已经用黄色标注出来。
// 
//
// 示例 2： 
//
// 
//输入：edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
//输出：23
// 
//
// 示例 3： 
//
// 
//输入：edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
//输出：1
//解释：节点 0 与图的其余部分没有连通，所以只有节点 0 可以到达。
// 
//
// 
//
// 提示： 
//
// 
// 0 <= edges.length <= min(n * (n - 1) / 2, 10⁴) 
// edges[i].length == 3 
// 0 <= ui < vi < n 
// 图中 不存在平行边 
// 0 <= cnti <= 10⁴ 
// 0 <= maxMoves <= 10⁹ 
// 1 <= n <= 3000 
// 
// Related Topics 图 最短路 堆（优先队列） 👍 40 👎 0


package leetcode

import (
	"container/heap"
)

//leetcode submit region begin(Prohibit modification and deletion)
func reachableNodes(edges [][]int, maxMoves int, n int) int {
	graph := make([][]info, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], info{v, w+1})
		graph[v] = append(graph[v], info{u, w+1})
	}
	visited := make([]bool, n)
	h := &infoHeap{}
	// 最短距离
	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1e9
	}
	dist[0] = 0
	heap.Init(h)
	heap.Push(h, info{0, 0})
	// dijkstra
	for h.Len() > 0 {
		nodeInfo := heap.Pop(h).(info)
		if visited[nodeInfo.nodeIndex] {
			continue
		}
		visited[nodeInfo.nodeIndex] = true
		for _, neighbor := range graph[nodeInfo.nodeIndex] {
			if nodeInfo.weight+neighbor.weight < dist[neighbor.nodeIndex] {
				dist[neighbor.nodeIndex] = nodeInfo.weight+neighbor.weight
				heap.Push(h, info{neighbor.nodeIndex, nodeInfo.weight+neighbor.weight})
			}
		}
	}
	var res int
	// 新增点
	for _, edge := range edges {
		left := max(0, min(maxMoves-dist[edge[0]], edge[2]))
		right := max(0, min(maxMoves-dist[edge[1]], edge[2]))
		res += min(left+right, edge[2])
	}
	// 原顶点
	for i := 0; i < n; i++ {
		if dist[i] <= maxMoves {
			res++
		}
	}
	return res
}

type info struct {
	nodeIndex int
	weight int
}

type infoHeap []info

func (this infoHeap)Len() int {
	return len(this)
}

func (this infoHeap)Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this infoHeap)Less(i, j int) bool {
	// 小顶堆
	if this[i].weight != this[j].weight {
		return this[i].weight < this[j].weight
	} else {
		return this[i].nodeIndex < this[j].nodeIndex
	}
}

func (this *infoHeap)Push(x interface{}) {
	*this = append(*this, x.(info))
}

func (this *infoHeap)Pop() interface{} {
	x := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return x
}
//leetcode submit region end(Prohibit modification and deletion)
