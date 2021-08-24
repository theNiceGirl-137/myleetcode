//给你一个数组 nums ，请你完成两类查询，其中一类查询要求更新数组下标对应的值，另一类查询要求返回数组中某个范围内元素的总和。 
//
// 实现 NumArray 类： 
//
// 
// 
// 
// NumArray(int[] nums) 用整数数组 nums 初始化对象 
// void update(int index, int val) 将 nums[index] 的值更新为 val 
// int sumRange(int left, int right) 返回子数组 nums[left, right] 的总和（即，nums[left] + 
//nums[left + 1], ..., nums[right]） 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["NumArray", "sumRange", "update", "sumRange"]
//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//输出：
//[null, 9, null, 8]
//
//解释：
//NumArray numArray = new NumArray([1, 3, 5]);
//numArray.sumRange(0, 2); // 返回 9 ，sum([1,3,5]) = 9
//numArray.update(1, 2);   // nums = [1,2,5]
//numArray.sumRange(0, 2); // 返回 8 ，sum([1,2,5]) = 8
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 3 * 104 
// -100 <= nums[i] <= 100 
// 0 <= index < nums.length 
// -100 <= val <= 100 
// 0 <= left <= right < nums.length 
// 最多调用 3 * 104 次 update 和 sumRange 方法 
// 
// 
// 
// Related Topics 设计 树状数组 线段树 数组 
// 👍 290 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	// 线段树
	n int
	tree []int
}


func Constructor307(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, 2*n)
	// 后半部分存储原数组的值
	for i, j := n, 0; i < len(tree); i, j = i+1, j+1 {
		tree[i] = nums[j]
	}
	// 前半部分表示非叶节点，存储区间的值
	// 节点 i 的左右子树分别是 2*i 和 2*i+1
	for i := n-1; i > 0; i-- {
		tree[i] = tree[2*i]+tree[2*i+1]
	}
	return NumArray{
		n: n,
		tree: tree,
	}
}


func (this *NumArray) Update(index int, val int)  {
	// 在线段树中原数组对应的位置
	index += this.n
	// 需要在原来的值上做出的更改
	val -= this.tree[index]
	for i := index; i > 0; i /= 2 {
		this.tree[i] += val
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	var ans int
	left, right = left+this.n, right+this.n
	for left <= right {
		//奇数表示右节点，偶数表示左节点,需要单独把偏移位置的值加上
		// 然后向上找到区间节点
		if left&1 == 1 {
			ans += this.tree[left]
			left++
		}
		if right&1 == 0 {
			ans += this.tree[right]
			right--
		}
		//变成父节点
		left, right = left/2, right/2
	}
	return ans
}
/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
