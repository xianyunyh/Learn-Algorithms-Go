package hash

//
func groupAnagrams(strs []string) [][]string {
	hashMap := map[int][]string{}
	keys := map[rune]int{
		'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19, 'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41,
		'n': 43, 'o': 47, 'p': 53, 'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89, 'y': 97, 'z': 101,
	}
	hashFn := func(str string) int {
		sum := 1
		for _, v := range str {
			sum *= keys[v]
		}
		return sum
	}

	for _, item := range strs {
		code := hashFn(item)
		if v, ok := hashMap[code]; ok && len(v) > 0 {
			hashMap[code] = append(hashMap[code], item)
		} else {
			hashMap[code] = []string{item}
		}
	}
	results := make([][]string, 0)
	for _, v := range hashMap {
		results = append(results, v)
	}
	return results
}
