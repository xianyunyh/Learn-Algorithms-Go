package hash

// 解题思路 就是利用map存储计数，重复数的计数肯定大于1
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
