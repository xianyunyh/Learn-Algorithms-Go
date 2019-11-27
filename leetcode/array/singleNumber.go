package array

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 利用map的映射。 然后key作为值， 计数作为值
func singleNumber(nums []int) int {
	tempMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := tempMap[v]; ok {
			tempMap[v] = tempMap[v] + 1
		} else {
			tempMap[v] = 1
		}
	}
	for v, n := range tempMap {
		if n == 1 {
			return v
		}
	}
	return 0
}
