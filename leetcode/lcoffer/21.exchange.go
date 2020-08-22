package lcoffer

func exchange(nums []int) []int {
	left, right := 0, (len(nums) - 1)
	for left < right {
		if nums[left]%2 != 0 {
			left++
			continue
		}
		if nums[right]%2 != 1 {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
