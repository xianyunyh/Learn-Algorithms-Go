package lcoffer

//利用map 当数值重复时直接返回 0（N）
func findRepeatNumber(nums []int) int {
	table := make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := table[v]; !ok {
			table[v] = 1
		} else {
			return v
		}
	}
	return -1
}

//数组当hash表。
func findRepeatNumber2(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		temp := nums[i]
		nums[i] = nums[temp]
		nums[temp] = temp
	}
	return -1
}
