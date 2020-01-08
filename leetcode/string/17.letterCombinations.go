package string

//
func letterCombinations(digits string) []string {
	table := map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}
	for _, v := range digits {
		n := int(v - '0')
		nums := table[n]
		pre := result
		result = []string{}
		for _, char := range nums {
			for _, v2 := range pre {
				result = append(result, v2+char)
			}
		}
	}
	return result
}
