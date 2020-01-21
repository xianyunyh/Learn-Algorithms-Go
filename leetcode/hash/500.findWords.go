package hash

func findWords(words []string) []string {
	keywords := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}
	table := make(map[rune]int)
	for i, keyword := range keywords {
		for _, c := range keyword {
			table[c] = i
			table[c-32] = i
		}
	}
	results := []string{}
	for _, word := range words {
		first := table[rune(word[0])]
		r := true
		for _, v := range word {
			if table[v] != first {
				r = false
			}
		}
		if r == true {
			results = append(results, word)
		}
	}
	return results
}
