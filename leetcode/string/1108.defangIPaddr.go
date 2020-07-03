package string

func defangIPaddr(address string) string {
	if len(address) <= 1 {
		return address
	}
	result := []byte{}
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			result = append(result, []byte("[.]")...)
		} else {
			result = append(result, address[i])
		}
	}
	return string(result)
}

func replace(str string, old byte, need string) string {
	if len(need) == 0 || len(str) == 0 {
		return str
	}
	result := []byte{}
	for i := 0; i < len(str); i++ {
		if str[i] == old {
			result = append(result, []byte(need)...)
		} else {
			result = append(result, str[i])
		}
	}
	return string(result)
}
