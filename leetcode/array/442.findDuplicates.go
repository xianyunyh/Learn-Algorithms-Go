package array

// [442] 数组中重复的数据 哈希表
func findDuplicates(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	result := make([]int, 0)
	for n, v := range m {
		if v == 2 {
			result = append(result, n)
		}
	}
	return result
}
