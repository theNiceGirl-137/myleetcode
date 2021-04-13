//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。 
//
// 注意: 
//
// 
// 可以认为区间的终点总是大于它的起点。 
// 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。 
// 
//
// 示例 1: 
//
// 
//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//
//输出: 1
//
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
// 
//
// 示例 2: 
//
// 
//输入: [ [1,2], [1,2], [1,2] ]
//
//输出: 2
//
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
// 
//
// 示例 3: 
//
// 
//输入: [ [1,2], [2,3] ]
//
//输出: 0
//
//解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
// 
// Related Topics 贪心算法 
// 👍 400 👎 0


package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}
	// 区间长度一定为 2，形式都形如[a, b]
	// 将区间按结尾大小增序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// 保留区间时，选择的区间的结尾越小，留给其它区间的空间就越大，就能保留更多的空间
	// 所以采取的贪心策略为优先保留结尾小且不相交的区间
	remove := 0
	pre := intervals[0][1]
	// 每次选择结尾最小且和前一个区间不重叠的区间
	for i := 1; i < len(intervals); i++ {
		// 后一个区间的开头小于前一个区间的结尾，一定相交
		if intervals[i][0] < pre {
			remove++
		} else {
			// 更新待比较区间
			pre = intervals[i][1]
		}
	}
	return remove
}
//leetcode submit region end(Prohibit modification and deletion)
