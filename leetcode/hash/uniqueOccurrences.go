package hash

// 给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
//
//如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false
//利用map 和 set 构建
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	set := make(map[int]struct{})
	for _,v := range arr {
		if _,ok := m[v];ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	for _,v := range m {
		if _,ok := set[v];ok {
			return false
		}
		set[v] = struct{}{}
	}
	return true
}
