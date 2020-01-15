package hash

func isAnagram(s string, t string) bool {
	table := [26]int{}
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		table[int(v-'a')]++
		table[int(t[i]-'a')]--
	}
	for _, v := range table {
		if v != 0 {
			return false
		}
	}
	return true
}
