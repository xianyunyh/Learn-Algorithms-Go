package string

func lengthOfLastWord(str string) int {
	strLen := len(str)
	if strLen == 0 {
		return 0
	}
	if str != " " && strLen == 1 {
		return 1
	}
	j := 0
	for i := strLen - 1; i >= 0; i-- {
		if string(str[i]) == " " {
			if j == 0 {
				continue
			}
			break
		} else {
			j++
		}
	}
	return j
}
