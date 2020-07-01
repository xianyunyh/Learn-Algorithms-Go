package string

// TODO
// 思路
// 用子串的第一个第一个字符串进行
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	max := len(haystack) - len(needle) + 1
	var i, j int
	for i = 0; i < max; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
