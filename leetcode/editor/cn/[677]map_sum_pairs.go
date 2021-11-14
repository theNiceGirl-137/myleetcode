//实现一个 MapSum 类，支持两个方法，insert 和 sum： 
//
// 
// MapSum() 初始化 MapSum 对象 
// void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 
//key 已经存在，那么原来的键值对将被替代成新的键值对。 
// int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["MapSum", "insert", "sum", "insert", "sum"]
//[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
//输出：
//[null, null, 3, null, 5]
//
//解释：
//MapSum mapSum = new MapSum();
//mapSum.insert("apple", 3);  
//mapSum.sum("ap");           // return 3 (apple = 3)
//mapSum.insert("app", 2);    
//mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)
// 
//
// 
//
// 提示： 
//
// 
// 1 <= key.length, prefix.length <= 50 
// key 和 prefix 仅由小写英文字母组成 
// 1 <= val <= 1000 
// 最多调用 50 次 insert 和 sum 
// 
// Related Topics 设计 字典树 哈希表 字符串 👍 163 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
type MapSum struct {
	children [26]*MapSum
	value int
	isLeaf bool
}


func Constructor() MapSum {
	return MapSum{}
}


func (this *MapSum) Insert(key string, val int)  {
	m := this
	for _, ch := range key {
		ch -= 'a'
		if m.children[ch] == nil {
			m.children[ch] = &MapSum{}
		}
		m = m.children[ch]
	}
	m.isLeaf = true
	m.value = val
}


func (this *MapSum) Sum(prefix string) int {
	m := this.SearchPrefix(prefix)
	sum := 0
	if m != nil {
		m.FindLeaf(&sum)
	}
	return sum
}

func (this *MapSum) SearchPrefix(prefix string) *MapSum {
	m := this
	for _, ch := range prefix {
		ch -= 'a'
		if m.children[ch] == nil {
			return nil
		}
		m = m.children[ch]
	}
	return m
}

func (this *MapSum) FindLeaf(sum *int) {
	if this.isLeaf == true {
		*sum += this.value
	}
	for i := 0; i < 26; i++ {
		if this.children[i] != nil {
			this.children[i].FindLeaf(sum)
		}
	}
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
