//给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。这个矩形的左下顶点是
// (xi, yi) ，右上顶点是 (ai, bi) 。 
//
// 如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。 
// 
//
// 示例 1： 
//
// 
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
//输出：true
//解释：5 个矩形一起可以精确地覆盖一个矩形区域。 
// 
//
// 示例 2： 
//
// 
//输入：rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
//输出：false
//解释：两个矩形之间有间隔，无法覆盖成一个矩形。 
//
// 示例 3： 
//
// 
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[3,2,4,4]]
//输出：false
//解释：图形顶端留有空缺，无法覆盖成一个矩形。 
//
// 示例 4： 
//
// 
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
//输出：false
//解释：因为中间有相交区域，虽然形成了矩形，但不是精确覆盖。 
//
// 
//
// 提示： 
//
// 
// 1 <= rectangles.length <= 2 * 10⁴ 
// rectangles[i].length == 4 
// -10⁵ <= xi, yi, ai, bi <= 10⁵ 
// 
// Related Topics 数组 扫描线 👍 146 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func isRectangleCover(rectangles [][]int) bool {
	// 左下角、左上角、右下角、右上角的顶点一定是所有矩阵的最外圈
	// 它们的面积和一定和完美矩阵的面积相等
	// 任意不是这四个顶点的小矩阵的点，一定会出现两次或四次（如果出现四次以上，一定有超过四个矩阵以这个点为顶点，那么必然有重叠；如果出现奇数次，那么必然没有被完整覆盖）
	x, y, a, b, s := 10001, 10001, -10001, -10001, 0
	cnts := map[int]int{}
	for _, r := range rectangles {
		x, y, a, b = min(x, r[0]), min(y, r[1]), max(a, r[2]), max(b, r[3])
		s += (r[2] - r[0]) * (r[3] - r[1])
		cnts[point(r[0], r[1])] += 1
		cnts[point(r[0], r[3])] += 1
		cnts[point(r[2], r[1])] += 1
		cnts[point(r[2], r[3])] += 1
	}
	if s != (a-x)*(b-y) {
		return false
	}
	points := map[int]bool{}
	points[point(x, y)] = true
	points[point(x, b)] = true
	points[point(a, y)] = true
	points[point(a, b)] = true
	for p := range points {
		v, err := cnts[p]
		if !err || v > 1 {
			return false
		}
	}
	for p, v := range cnts {
		if !points[p] {
			if v > 4 || v%2 == 1 {
				return false
			}
		}
	}
	return true
}

func point(a, b int) int {
	return 10001 * a + b
}
//leetcode submit region end(Prohibit modification and deletion)
