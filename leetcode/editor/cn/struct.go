package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack
func top(arr []int) int {
	return arr[len(arr)-1]
}

func pop(arr []int) []int {
	return arr[:len(arr)-1]
}

func push(arr []int, x int) []int {
	return append(arr, x)
}

func sTop(arr []byte) byte {
	return arr[len(arr)-1]
}

func sPop(arr []byte) []byte {
	return arr[:len(arr)-1]
}

func sPush(arr []byte, x byte) []byte {
	return append(arr, x)
}

// Queue
type queue []int

func (this *queue)push(x int) {
	*this = append(*this, x)
}

func (this *queue)pop() {
	*this = (*this)[1:]
}

func (this *queue)empty() bool {
	return len(*this) == 0
}

func (this *queue)front() int {
	return (*this)[0]
}

// Heap
// 优先队列(priority queue)可以在 O(1) 时间内获得最大值，并且可以在 O(log n) 时间内取出最大值或插入任意值
// 优先队列常常用堆(heap)来实现，堆是一个完全二叉树，其每个节点的值总是大于等于子节点的值
// 实际实现堆时，我们通常用一个数组而不是用指针建立一个树
// 这是因为堆是完全二叉树，所以用数组表示时，位置 i 的节点的父节点位置一定为 i/2，而它的两个子节点的位置又一定分别为 2i 和 2i+1
// 以下是堆的实现方法，其中最核心的两个操作是上浮和下沉：如果一个节点比父节点大，那么需要交换这个两个节点
// 交换后还可能比它新的父节点大，因此需要不断地进行比较和交换操作，我们称之为上浮
// 类似地，如果一个节点比父节小，也需要不断地向下进行比较和交换操作，我们称之为下沉
// 如果一个节点有两个子节点，我们总是交换最大的子节点
// 获取最大值
func hTop(heap []int) int {
	return heap[0]
}

// 插入任意值，把新的数字放在最后一位，然后上浮
func hPush(k int, heap []int) []int {
	heap = append(heap, k)
	swim(len(heap)-1, heap)
	return heap
}

// 删除最大值，把最后一个数字挪到开头，然后下沉
func hPop(heap []int) []int {
	heap[0] = heap[len(heap)-1]
	heap = heap[1:]
	sink(0, heap)
	return heap
}

// 上浮
func swim(pos int, heap []int) {
	for pos > 1 && heap[pos/2] < heap[pos] {
		heap[pos/2], heap[pos] = heap[pos], heap[pos/2]
		pos /= 2
	}
}

// 下沉
func sink(pos int, heap []int) {
	for 2*pos <= len(heap) {
		i := 2*pos
		if i < len(heap) && heap[i] < heap[i+1] {
			i++
		}
		if heap[pos] >= heap[i] {
			break
		}
		heap[pos], heap[i] = heap[i], heap[pos]
		pos = i
	}
}
// 通过将算法中的大于号和小于号互换，我们也可以得到一个快速获得最小值的优先队列
// 如果我们需要在维持大小关系的同时，还需要支持查找任意值、删除任意值、维护所有数字的大小关系等操作，可以考虑使用 set 或 map 来代替优先队列

// BST
func makeEmpty(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	makeEmpty(t.Left)
	makeEmpty(t.Right)
	return new(TreeNode)
}

func insert(t *TreeNode, x int) *TreeNode {
	if t == nil {
		t = new(TreeNode)
		t.Val = x
		t.Left, t.Right = nil, nil
	} else if x < t.Val {
		t.Left = insert(t.Left, x)
	} else if x > t.Val {
		t.Right = insert(t.Right, x)
	}
	return t
}

func find(t *TreeNode, x int) *TreeNode {
	if t == nil {
		return nil
	}
	if x < t.Val {
		return find(t.Left, x)
	}
	if x > t.Val {
		return find(t.Right, x)
	}
	return t
}

func findMin(t *TreeNode) *TreeNode {
	if t == nil || t.Left == nil {
		return t
	}
	return findMin(t.Left)
}

func findMax(t *TreeNode) *TreeNode {
	if t == nil || t.Right == nil {
		return t
	}
	return findMin(t.Right)
}

func remove(t *TreeNode, x int) *TreeNode {
	temp := new(TreeNode)
	if t == nil {
		return nil
	} else if x < t.Val {
		t.Left = remove(t.Left, x)
	} else if x > t.Val {
		t.Right = remove(t.Right, x)
	} else if t.Left != nil && t.Right != nil {
		temp = findMin(t.Right)
		t.Val = temp.Val
		t.Right = remove(t.Right, t.Val)
	} else {
		if t.Left == nil {
			t = t.Right
		} else if t.Right == nil {
			t = t.Left
		}
	}
	return t
}