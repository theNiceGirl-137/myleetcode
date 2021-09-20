package leetcode

// 工具函数
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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// 反转数组
func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

// 反转字符串
func reverseStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Stack tool
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

