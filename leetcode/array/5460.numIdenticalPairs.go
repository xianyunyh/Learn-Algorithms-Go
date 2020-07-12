package array

func numIdenticalPairs(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		for j := 1 + i; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				n += 1
			}
		}
	}
	return n
}
