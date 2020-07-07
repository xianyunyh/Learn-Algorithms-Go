package math

func singleNumber(nums []int) int {
	one, two, three := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		two |= (one & nums[i])
		one ^= nums[i]
		three = (one & two)
		one ^= three
		two ^= three
	}
	return one
}
