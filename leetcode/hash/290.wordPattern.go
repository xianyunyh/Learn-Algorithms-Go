package hash

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	table := make(map[rune]string)
	set := make(map[string]rune)

	for k, v := range pattern {
		if _, ok := table[v]; !ok {
			table[v] = strs[k]
			if _, ok2 := set[strs[k]]; !ok2 {
				set[strs[k]] = v
			} else {
				if set[strs[k]] != v {
					return false
				}
			}
		} else {
			if table[v] != strs[k] {
				return false
			}
		}
	}
	return true
}
