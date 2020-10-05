package array

// 寻找中心索引， 2*T + nums[i] = Sum
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	sum2 := 0
	for i := 0; i < len(nums); i++ {
		if sum-nums[i]-sum2 == sum2 {
			return i
		}
		sum2 += nums[i]
	}
	return -1
}
