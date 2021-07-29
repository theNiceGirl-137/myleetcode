//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
// 
//
// 示例 2: 
//
// 
//输入: nums = [1], k = 1
//输出: [1] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 105 
// k 的取值范围是 [1, 数组中不相同的元素的个数] 
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的 
// 
//
// 
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。 
// Related Topics 堆 哈希表 
// 👍 732 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	maxCount := 0
	// 为每个值设立一个桶，桶内统计这个值出现的次数
	for i := range nums {
		counts[nums[i]]++
		maxCount = max(maxCount, counts[nums[i]])
	}
	// 对桶的频次进行排序
	buckets := make([][]int, maxCount+1)
	res := make([]int, 0)
	// 按出现的次数将数字放进新桶内
	// 出现次数越大，即 v 值更靠后，就是出现频率更高的值
	for k, v := range counts {
		buckets[v] = append(buckets[v], k)
	}
	// 从后往前遍历，找到 k 个旧桶
	for i := len(buckets)-1; i > 0; i-- {
		if len(buckets[i]) > 0 {
			for j := range buckets[i] {
				res = append(res, buckets[i][j])
				if len(res) == k {
					return res
				}
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
