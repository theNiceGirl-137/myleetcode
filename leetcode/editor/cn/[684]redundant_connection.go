//树可以看成是一个连通且 无环 的 无向 图。 
//
// 给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信
//息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。 
//
// 请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入: edges = [[1,2], [1,3], [2,3]]
//输出: [2,3]
// 
//
// 示例 2： 
//
// 
//
// 
//输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
//输出: [1,4]
// 
//
// 
//
// 提示: 
//
// 
// n == edges.length 
// 3 <= n <= 1000 
// edges[i].length == 2 
// 1 <= ai < bi <= edges.length 
// ai != bi 
// edges 中无重复元素 
// 给定的图是连通的 
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 392 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func findRedundantConnection(edges [][]int) []int {
	// 判断两个结点是否被重复连接，使用并查集
	id := make([]int, len(edges)+1)
	size := make([]int, len(edges)+1)
	for i := 0; i < len(id); i++ {
		id[i] = i
		size[i] = 1
	}
	uf := UF{id: id, size: size}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if uf.isConnected(u, v) {
			return edge
		}
		uf.connect(u, v)
	}
	return []int{-1, -1}
}

type UF struct {
	id []int
	size []int
}

func (this UF)find(p int) int {
	//for p != this.id[p] {
	//	p = this.id[p]
	//}
	// 路径压缩，使得下次查找更快
	for p != this.id[p] {
		this.id[p] = this.id[this.id[p]]
		p = this.id[p]
	}
	return p
}

func (this UF)connect(p, q int) {
	//this.id[this.find(p)] = this.find(q)
	// 按轶合并：每次合并都把深度较小的集合合并在深度较大的集合下面
	i, j := this.find(p), this.find(q)
	if i != j {
		if this.size[i] < this.size[j] {
			this.id[i] = j
			this.size[j] += this.size[i]
		} else {
			this.id[j] = i
			this.size[i] += this.size[j]
		}
	}
}

func (this UF)isConnected(p, q int) bool {
	return this.find(p) == this.find(q)
}
//leetcode submit region end(Prohibit modification and deletion)
