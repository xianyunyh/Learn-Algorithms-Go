package array

func maxProduct(nums []int) int {
	max := nums[0]
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
			maxIndex = i
		}
	}
	second := 0
	for i := 0; i < len(nums); i++ {
		if second < nums[i] && nums[i] <= max && i != maxIndex {
			second = nums[i]
		}
	}
	return (max - 1) * (second - 1)
}
