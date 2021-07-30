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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

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

