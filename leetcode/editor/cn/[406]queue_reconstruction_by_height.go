//假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 
//个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。 
//
// 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第
// j 个人的属性（queue[0] 是排在队列前面的人）。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
//输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
//解释：
//编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
//编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
//编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
//编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
//编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
//编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
//因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
// 
//
// 示例 2： 
//
// 
//输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
//输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= people.length <= 2000 
// 0 <= hi <= 106 
// 0 <= ki < people.length 
// 题目数据确保队列可以被重建 
// 
// Related Topics 贪心算法 
// 👍 848 👎 0


package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func reconstructQueue(people [][]int) [][]int {
	// 先按照身高将队列排序，在身高相等的情况下按照 k 值的大小排序
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	// 按照 k 值插入到 index=k 的地方，index 之后的往后移动
	// 排序过后，p[1] 就是这个人应该站的位置，因为高个子的人在排序过后的顺序是正确的，需要将矮个子的人
	// 按照他们应该站的位置插入到高个子的人之间
	for i, p := range people {
		// 把 p[1]~i 位置上的人移动到 p[1]+1~i 位置上，将 p[1] 位置填充上p
		// 如果 p[1] 就在 i 处是不会进行 copy 的，p 就在 i 处，copy 之后不用担心覆盖问题，p 的值已被记录
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p
	}
	return people
}
//leetcode submit region end(Prohibit modification and deletion)
