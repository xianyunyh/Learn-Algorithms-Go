package string

// TODO
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	i := 0
	j := 0
	char := needle[i]
	for ; j < len(haystack); j++ {
		if haystack[j] != char {
			continue
		}
		i++
		if i == len(needle) {
			break
		}
		char = needle[i]
	}
	if i == len(needle) {
		return j - len(needle) + 1
	}
	return -1

}
