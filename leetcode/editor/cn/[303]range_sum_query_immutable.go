//给定一个整数数组 nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。 
//
// 
// 
// 实现 NumArray 类： 
//
// 
// NumArray(int[] nums) 使用数组 nums 初始化对象 
// int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 s
//um(nums[i], nums[i + 1], ... , nums[j])） 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["NumArray", "sumRange", "sumRange", "sumRange"]
//[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
//输出：
//[null, 1, -1, -3]
//
//解释：
//NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
//numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
//numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1)) 
//numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 104 
// -105 <= nums[i] <= 105 
// 0 <= i <= j < nums.length 
// 最多调用 104 次 sumRange 方法 
// 
// 
// 
// Related Topics 设计 数组 前缀和 
// 👍 355 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
type NumArray303 struct {
	psum []int
}


func Constructor303(nums []int) NumArray303 {
	// 对于一维的数组，我们可以使用前缀和来解决此类问题。先建立一个与数组 nums 长度相
	// 同的新数组 psum，表示 nums 每个位置之前前所有数字的和
	// 可以直接遍历一遍 nums 数组，并利用状态转移方程 psum[i] = psum[i-1] + nums[i] 完成统计
	// 如果我们需要获得位置 i 和 j 之间的数字和，只需计算 psum[j+1] - psum[i] 即可
	numArray := NumArray303{
		psum: make([]int, len(nums)+1),
	}
	numArray.psum[0] = 0
	for i := 0; i < len(nums); i++ {
		numArray.psum[i+1] = numArray.psum[i]+nums[i]
	}
	return numArray
}


func (this *NumArray303) SumRange(left int, right int) int {
	return this.psum[right+1]-this.psum[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
