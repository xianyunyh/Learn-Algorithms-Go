package array

// [9] 回文数
func isPalindrome(x int) bool {
	arr := make([]int, 0)
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	for x >= 10 {
		arr = append(arr, x%10)
		x = x / 10

	}
	arr = append(arr, x)

	lens := len(arr)
	for i := 0; i < lens; i++ {
		if arr[i] != arr[lens-1-i] {
			return false
		}
	}
	return true
}
