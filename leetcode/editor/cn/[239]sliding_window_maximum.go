//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。 
//
// 返回滑动窗口中的最大值。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 
//
// 示例 2： 
//
// 
//输入：nums = [1], k = 1
//输出：[1]
// 
//
// 示例 3： 
//
// 
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
// 
//
// 示例 4： 
//
// 
//输入：nums = [9,11], k = 2
//输出：[11]
// 
//
// 示例 5： 
//
// 
//输入：nums = [4,-2], k = 2
//输出：[4] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 105 
// -104 <= nums[i] <= 104 
// 1 <= k <= nums.length 
// 
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 
// 👍 1129 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	// 我们可以利用双端队列进行操作：每当向右移动时，把窗口左端的值从队列左端剔除，把队
	// 列右边小于窗口右端的值全部剔除,这样双端队列的最左端永远是当前窗口内的最大值
	// 另外，这道题也是单调栈的一种延伸：该双端队列利用从左到右递减来维持大小关系
	dq := Deque{queue: make([]int, 0)}
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		dq.Push(nums[i])
	}
	ans = append(ans, dq.Front())
	for i := k ; i < len(nums); i++ {
		// 每次移除最前面的元素
		dq.Pop(nums[i-k])
		// 添加新的元素
		dq.Push(nums[i])
		ans = append(ans, dq.Front())
	}
	return ans
}

type Deque struct {
	queue []int
}

func (dq *Deque) Front() int {
	return dq.queue[0]
}

func (dq *Deque) Back() int {
	return dq.queue[len(dq.queue)-1]
}

func (dq *Deque) Empty() bool {
	return len(dq.queue) == 0
}

func (dq *Deque) Push(x int) {
	for !dq.Empty() && x > dq.Back() {
		dq.queue = dq.queue[:len(dq.queue)-1]
	}
	dq.queue = append(dq.queue, x)
}

func (dq *Deque) Pop(x int) {
	if !dq.Empty() && x == dq.Front() {
		dq.queue = dq.queue[1:]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
