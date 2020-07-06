package hash

// 同387
func FirstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	table := make(map[rune]int)
	index := make(map[rune]int)

	for k, v := range s {
		if _, ok := index[v]; !ok {
			index[v] = k
		}
		table[v] += 1
	}
	j := len(s)
	for v, n := range table {
		if n == 1 {
			idx := index[v]
			if idx <= j {
				j = idx
			}
		}
	}
	if j == len(s) {
		return ' '
	}
	return s[j]
}
