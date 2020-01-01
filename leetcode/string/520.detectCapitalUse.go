package string

func detectCapitalUse(word string) bool {

	first := false
	upperLetters := 0
	isUpper := func(v rune) bool {
		return v >= 65 && v <= 90
	}
	for i, v := range word {
		if i == 0 && (isUpper(v)) {
			first = true
			upperLetters++
			continue
		}
		if isUpper(v) {
			upperLetters++
		}
	}
	//单词中所有字母都不是大写
	if upperLetters == 0 {
		return true
	}
	// 全部字母都是大写
	if upperLetters == len(word) {
		return true
	}
	//如果单词不只含有一个字母，只有首字母大写
	if upperLetters == 1 && first {
		return true
	}
	return false
}
