package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(result) && j < len(strs[i]); j++ {
			if result[j] != strs[i][j] {
				break
			}
		}
		result = result[0:j]
		if result == "" {
			return result
		}
	}
	return result
}
