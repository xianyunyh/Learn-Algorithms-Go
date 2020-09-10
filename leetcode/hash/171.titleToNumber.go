package hash

func titleToNumber(s string) int {
	res := 0
	for _, v := range s {
		num := int(v-'A') + 1
		res = res*26 + num
	}
	return res
}
