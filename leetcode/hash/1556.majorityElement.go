package hash

func majorityElement(nums []int) int {
	table := make(map[int]int)

	for _, v := range nums {
		table[v]++
	}
	for k, v := range table {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}
