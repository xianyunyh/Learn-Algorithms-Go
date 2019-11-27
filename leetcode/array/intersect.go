package array

//计算两个数组之间的交集
//利用hash表

func Intersect(a, b []int) []int {
	result := []int{}
	mapA := make(map[int]int)
	mapB := make(map[int]int)

	for _, v := range a {
		if _, ok := mapA[v]; ok {
			mapA[v] += 1
		} else {
			mapA[v] += 1
		}
	}
	for _, v := range b {
		if _, ok := mapB[v]; ok {
			mapB[v] += 1
		} else {
			mapB[v] = 1
		}

	}
	for k, v := range mapA {
		num, ok := mapB[k]
		if !ok {
			continue
		}

		if v <= 1 || num <= 1 {
			result = append(result, k)
			continue
		}
		max := k
		if v < max {
			max = v
		}
		for max > 0 {
			result = append(result, k)
			max--
		}
	}
	return result
}
