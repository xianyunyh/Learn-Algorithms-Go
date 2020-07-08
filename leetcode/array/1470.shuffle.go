package array

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	j := 0
	for i := 0; i < n; i++ {
		res[j] = nums[i]
		res[j+1] = nums[n+i]
		j += 2
	}
	return res
}
