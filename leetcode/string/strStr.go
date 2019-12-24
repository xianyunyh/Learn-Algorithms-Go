package string

// TODO
// 思路
// 用子串的第一个第一个字符串进行
func strStr(haystack string, needle string) int {
	//return strings.Index(haystack, needle)
	strLen := len(haystack)
	subStrLen := len(needle)
	if subStrLen == 0 {
		return 0
	}
	if strLen < subStrLen {
		return -1
	}
	i := 0
	j := 0
	for ; j < strLen; j++ {
		if haystack[j] != needle[i] {
			if i > 0 {
				j = j - i
				i = 0
			}
			continue
		}
		i++
		if i == len(needle) {
			break
		}
	}
	if i == len(needle) {
		return j - len(needle) + 1
	}
	return -1
}
