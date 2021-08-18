//城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回由这些建筑物形成的 天际线 。 
//
// 每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti] 表示： 
//
//
// 
// lefti 是第 i 座建筑物左边缘的 x 坐标。 
// righti 是第 i 座建筑物右边缘的 x 坐标。 
// heighti 是第 i 座建筑物的高度。 
// 
//
// 天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。关键点是水平线段的左端点。
//列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。 
//
// 注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答
//案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...] 
//
// 
//
// 示例 1： 
//
// 
//输入：buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
//输出：[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
//解释：
//图 A 显示输入的所有建筑物的位置和高度，
//图 B 显示由这些建筑物形成的天际线。图 B 中的红点表示输出列表中的关键点。 
//
// 示例 2： 
//
// 
//输入：buildings = [[0,2,3],[2,5,3]]
//输出：[[0,3],[5,0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= buildings.length <= 104 
// 0 <= lefti < righti <= 231 - 1 
// 1 <= heighti <= 231 - 1 
// buildings 按 lefti 非递减排序 
// 
// Related Topics 树状数组 线段树 数组 分治 有序集合 扫描线 堆（优先队列） 
// 👍 526 👎 0


package leetcode

import (
	"container/heap"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func getSkyline(buildings [][]int) [][]int {
	ans := make([][]int, 0)
	pointList := make([]*Point, 0)
	for _, building := range buildings {
		// 高度为负表示是从下到上，是左边高度的转折点
		pointList = append(pointList, &Point{p: building[0], h: -building[2]})
		pointList = append(pointList, &Point{p: building[1], h: building[2]})
	}
	sort.Slice(pointList, func(i, j int) bool {
		if pointList[i].p != pointList[j].p {
			return pointList[i].p < pointList[j].p
		}
		return pointList[i].h < pointList[j].h
	})
	// pre 为之前的高度，方便计算高度差，mapper 的 value 为需要删除的次数
	pre, mapper := 0, map[int]int{}
	maxHeaiPointPriority := &MaxPriority{sort.IntSlice{0}}
	heap.Init(maxHeaiPointPriority)
	for _, point := range pointList {
		if point.h < 0 {
			// 说明是从下到上，其高度可以参与关键点选择
			heap.Push(maxHeaiPointPriority, -point.h)
		} else {
			// 方块右端点计数，即跳过对应方块左端点的个数
			mapper[point.h]++
		}
		// 延迟删除
		for maxHeaiPointPriority.Len() > 0 {
			tmp := maxHeaiPointPriority.Head()
			if _, ok := mapper[tmp]; !ok {
				break
			}
			if mapper[tmp] == 1 {
				delete(mapper, tmp)
			} else {
				mapper[tmp]--
			}
			heap.Pop(maxHeaiPointPriority)
		}

		current := maxHeaiPointPriority.Head()
		// 有高度差，才有关键点出现
		if current != pre {
			ans = append(ans, []int{point.p, current})
			pre = current
		}
	}
	return ans
}

type Point struct {
	p, h int
}

type MaxPriority struct {
	sort.IntSlice
}

func (p *MaxPriority) Less(i, j int) bool {
	return (*p).IntSlice[i] > (*p).IntSlice[j]
}

func (p *MaxPriority) Push(x interface{}) {
	(*p).IntSlice = append((*p).IntSlice, x.(int))
}

func (p *MaxPriority) Pop() interface{} {
	v := (*p).IntSlice[(*p).Len()-1]
	(*p).IntSlice = (*p).IntSlice[:p.Len()-1]
	return v
}

func (p *MaxPriority) Head() int {
	return (*p).IntSlice[0]
}
//leetcode submit region end(Prohibit modification and deletion)
