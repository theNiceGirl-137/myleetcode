//给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。 
//
// 
//
// 示例 1： 
//
// 
//输入：points = [[1,1],[2,2],[3,3]]
//输出：3
// 
//
// 示例 2： 
//
// 
//输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
//输出：4
// 
//
// 
//
// 提示： 
//
// 
// 1 <= points.length <= 300 
// points[i].length == 2 
// -104 <= xi, yi <= 104 
// points 中的所有点 互不相同 
// 
// Related Topics 几何 哈希表 数学 
// 👍 338 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maxPoints(points [][]int) int {
	// 对于每个点，我们对其它点建立哈希表，统计同一斜率的点一共有多少个。这里利用的原理
	// 是，一条线可以由一个点和斜率而唯一确定。另外也要考虑斜率不存在和重复坐标的情况。
	// 本题也利用了一个小技巧：在遍历每个点时，对于数组中位置 i 的点，我们只需要考虑 i 之
	// 后的点即可，因为 i 之前的点已经考虑过 i 了。
	var maxCount, same, sameY int
	var dx, dy float64
	for i := 0; i < len(points); i++ {
		mapper := make(map[float64]int)
		same, sameY = 1, 1
		for j := i+1; j < len(points); j++ {
			if points[i][1] == points[j][1] {
				sameY++
				if points[i][0] == points[j][0] {
					same++
				}
			} else {
				dx = float64(points[i][0] - points[j][0])
				dy = float64(points[i][1] - points[j][1])
				mapper[(dx/dy)]++
			}
		}
		maxCount = max(maxCount, sameY)
		for _, value := range mapper {
			maxCount = max(maxCount, same+value)
		}
	}
	return maxCount
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
