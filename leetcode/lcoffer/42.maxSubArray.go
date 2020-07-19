package lcoffer

// dp 将每次的结果存入，然后用上一个值和0 .
func maxSubArray(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += max(nums[i-1], 0)
		sum = max(sum, nums[i])
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
