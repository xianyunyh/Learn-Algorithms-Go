package hash

// 解题思路 利用每个hash 计数
func intersect(nums1 []int, nums2 []int) []int {
	table := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := table[v]; ok {
			table[v]++
		} else {
			table[v] = 1
		}
	}
	result := []int{}
	for _, v := range nums2 {
		if _, ok := table[v]; ok && table[v] > 0 {
			result = append(result, v)
			table[v]--
		}
	}
	return result
}
