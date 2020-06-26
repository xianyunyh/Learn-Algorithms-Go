package array

func wordBreak(s string, wordDict []string) bool {
	str := ""
	lens := len(s)
	for i := 0; i < lens; i++ {
		str += string(s[i])
		if inArray(str, wordDict) {
			str = ""
		}
	}
	return str == ""
}

func inArray(str string, wordDict []string) bool {
	for _, v := range wordDict {
		if v == str {
			return true
		}
	}
	return false
}
