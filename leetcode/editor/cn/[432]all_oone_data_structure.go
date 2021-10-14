//请你实现一个数据结构支持以下操作： 
//
// 
// Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。 
// Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否则使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任
//何事情。key 保证不为空字符串。 
// GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串"" 。 
// GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。 
// 
//
// 
//
// 挑战： 
//
// 你能够以 O(1) 的时间复杂度实现所有操作吗？ 
// Related Topics 设计 哈希表 链表 双向链表 👍 104 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
type AllOne struct {
	mkv map[string]int
	mvk map[int]map[string]bool
	mvn map[int]*node
	lv *list
}

func Constructor432() AllOne {
	return AllOne{
		mkv: make(map[string]int),
		mvk: make(map[int]map[string]bool),
		mvn: make(map[int]*node),
		lv: &list{},
	}
}

func (this *AllOne) AddKV(key string, val int, onInc bool) {
	if m, ok := this.mvk[val]; ok {
		m[key] = true
	} else {
		this.mvk[val] = map[string]bool{key: true}
	}
	if len(this.mvk[val]) == 1 {
		n := &node{val: val}
		this.mvn[val] = n
		if onInc {
			if val != 1 {
				this.lv.insertAfter(n, this.mvn[val-1])
			} else {
				this.lv.pushFront(n)
			}
		} else {
			this.lv.insertBefore(n, this.mvn[val+1])
		}
	}
}

func (this *AllOne) RemoveKV(key string, val int) {
	m := this.mvk[val]
	if len(m) != 1 {
		delete(m, key)
	} else {
		delete(this.mvk, val)
		this.lv.unlink(this.mvn[val])
		delete(this.mvn, val)
	}
}

func (this *AllOne) GetOneKey(n *node) string {
	if n == nil {
		return ""
	}
	keys := this.mvk[n.val]
	for key := range keys {
		return key
	}
	return ""
}

func (this *AllOne) Inc(key string)  {
	this.mkv[key]++
	v := this.mkv[key]
	this.AddKV(key, v, true)
	if v != 1 {
		this.RemoveKV(key, v-1)
	}
}


func (this *AllOne) Dec(key string) {
	if _, ok := this.mkv[key]; !ok{
		return
	}
 	this.mkv[key]--
 	v := this.mkv[key]
 	if v != 0 {
 		this.AddKV(key, v, false)
	}
	this.RemoveKV(key, v+1)
}


func (this *AllOne) GetMaxKey() string {
	return this.GetOneKey(this.lv.tail)
}


func (this *AllOne) GetMinKey() string {
	return this.GetOneKey(this.lv.head)
}

type node struct {
	val int
	prev, next *node
}

type list struct {
	head, tail *node
}

func (this *list) pushFront(n *node) {
	n.prev = nil
	if this.head != nil {
		this.head.prev = n
	} else {
		this.tail = n
	}
	n.next = this.head
	this.head = n
}

func (this *list) insertBefore(n, before *node) {
	if before != this.head {
		n.prev, n.next = before.prev, before
		n.prev.next, before.prev = n, n
	} else {
		n.prev, n.next = nil, before
		before.prev, this.head = n, n
	}
}

func (this *list) insertAfter(n, after *node) {
	if after != this.tail {
		n.prev, n.next = after, after.next
		n.next.prev, after.next = n, n
	} else {
		n.prev, n.next = after, nil
		after.next, this.tail = n, n
	}
}

func (this *list) popFront() *node {
	if this.head == nil {
		return nil
	}
	h := this.head
	if h.next != nil {
		this.head = h.next
	} else {
		this.head, this.tail = nil, nil
	}
	return h
}

func (this *list) popBack() *node {
	if this.tail == nil {
		return nil
	}
	t := this.tail
	if t.prev != nil {
		this.tail = t.prev
	} else {
		this.head, this.tail = nil, nil
	}
	return t
}

func (this *list) unlink(n *node) {
	if this.head == n {
		this.popFront()
		return
	} else if this.tail == n {
		this.popBack()
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
//leetcode submit region end(Prohibit modification and deletion)
