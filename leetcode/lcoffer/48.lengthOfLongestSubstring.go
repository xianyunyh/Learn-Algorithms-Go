package lcoffer

// 最长不含重复字符的子字符串
// 利用一个hash表存储每个字符出现的位置，如果第二次出现，则计算两个
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	i := -1
	table := make(map[byte]int)
	for j := 0; j < len(s); j++ {
		c := s[j]
		if _, ok := table[c]; ok {
			i = max(table[c], i)
		}
		table[c] = j
		res = max(res, j-i)
	}
	return res
}
