package array

// 给定一个整数数组，判断是否存在重复元素。
//
//如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

func containsDuplicate(nums []int) bool {
	tempMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := tempMap[v]; ok {
			tempMap[v] = tempMap[v] + 1
		} else {
			tempMap[v] = 1
		}
	}
	for _, v := range tempMap {
		if v >= 2 {
			return true
		}
	}
	return false
}
