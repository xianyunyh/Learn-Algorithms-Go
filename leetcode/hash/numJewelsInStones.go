package hash

import (
	"strings"
)

//给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
//J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。
// O(m+n)
func NumJewelsInStones1(J string, S string) int {
	num := 0
	for _, v := range S {
		if strings.Contains(J, string(v)) {
			num++
		}
	}
	return num
}

//利用map
func NumJewelsInStones(J string, S string) int {
	num := 0
	m := make(map[rune]struct{})

	for _, v := range J {
		m[v] = struct{}{}
	}
	for _, v := range S {
		if _, ok := m[v]; ok {
			num = num + 1
		}
	}
	return num
}

// 65-90 97-122
func countSegments(s string) int {
	num := 0
	for i := 1; i < len(s); i++ {
		if s[i] == 32 && s[i-1] != 32 {
			num = num + 1
		}
	}
	if s[len(s)-1] != 32 {
		num += 1
	}
	return num
}
