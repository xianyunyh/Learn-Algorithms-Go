package hash

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	table := make(map[rune]int)

	for _, v := range s {
		if _, ok := table[v]; !ok {
			table[v] = 1
		} else {
			table[v] += 1
		}
	}

	for i, v := range s {
		if l, ok := table[v]; ok && l == 1 {
			return i
		}
	}
	return -1
}
