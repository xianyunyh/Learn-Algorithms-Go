package lcoffer

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	i, j := len(s)-1, len(s)-1
	str := []byte{}
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		str = append(str, s[i+1:j+1]...)
		str = append(str, ' ')
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}
	return strings.TrimSpace(string(str))
}
