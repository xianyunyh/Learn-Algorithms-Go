package lcoffer

import (
	"strings"
)

//利用库函数
func replaceSpace(s string) string {
	if len(s) <= 1 {
		return s
	}
	return strings.ReplaceAll(s, " ", "%20")
}

//遍历替换
func replaceSpace2(s string) string {
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
