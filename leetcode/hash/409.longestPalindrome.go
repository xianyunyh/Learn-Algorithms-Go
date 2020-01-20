package hash

func longestPalindrome(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	table := make(map[rune]int)
	for _, v := range s {
		if _, ok := table[v]; ok {
			table[v]++
		} else {
			table[v] = 1
		}
	}
	num := 0
	for _, v := range table {
		num += v / 2 * 2
		//这是计算是否含有单字符
		if v%2 == 1 && num%2 == 0 {
			num++
		}
	}
	return num
}
