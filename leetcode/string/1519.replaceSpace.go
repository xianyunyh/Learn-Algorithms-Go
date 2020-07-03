package string

func replaceSpace(s string) string {
	if len(s) <= 1 {
		return s
	}
	result := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
