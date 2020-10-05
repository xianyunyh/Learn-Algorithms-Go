package array

func findMaxConsecutiveOnes(nums []int) int {

	count := 0
	total := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, v := range nums {
		if v > 0 {
			count += 1
		} else {
			total = max(total, count)
			count = 0
		}
	}
	return max(total, count)
}
