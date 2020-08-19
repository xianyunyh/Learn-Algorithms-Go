package lcoffer

// 1. 找出n+1位的最大数，依次递减
func printNumbers(n int) []int {
	max := 1
	for n > 0 {
		max = max * 10
		n = n - 1
	}
	result := make([]int, max-1)
	for i := 0; i < max-1; i++ {
		result[i] = i + 1
	}
	return result
}
