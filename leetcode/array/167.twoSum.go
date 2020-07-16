package array

func twoSum2(numbers []int, target int) []int {

	max := len(numbers) - 1
	min := 0
	for min < max {
		sum := numbers[min] + numbers[max]
		if sum == target {
			return []int{min + 1, max + 1}
		}
		if sum > target {
			max--
		} else {
			min++
		}
	}
	return []int{-1, -1}
}
