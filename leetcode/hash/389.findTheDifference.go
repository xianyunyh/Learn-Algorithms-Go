package hash

func findTheDifference(s string, t string) byte {
	sum1 := 0
	sum2 := 0
	for _, v := range s {
		sum1 += int(v)
	}
	for _, v := range t {
		sum2 += int(v)
	}
	return byte(sum2 - sum1)
	// return '0'
}

// 哈希计数
func findTheDifference2(s string, t string) byte {
	table1 := make(map[rune]int)

	for _, v := range s {
		if _, ok := table1[v]; ok {
			table1[v]++
		} else {
			table1[v] = 1
		}
	}
	for _, v := range t {
		if _, ok := table1[v]; ok {
			table1[v]++
		} else {
			table1[v] = 1
		}
	}
	for k, v := range table1 {
		if v%2 == 1 {
			return byte(k)
		}
	}
	return '0'
	// return '0'
}
