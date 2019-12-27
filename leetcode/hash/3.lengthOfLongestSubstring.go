package hash

import "math"

func lengthOfLongestSubstring(str string) int {
	l := len(str)
	sum := 0
	i := 0
	j := 0
	m := make(map[uint8]int)
	for ; j < l; j++ {
		if n, ok := m[str[j]]; ok {
			i = int(math.Max(float64(j), float64(n)))
		}
		sum = int(math.Max(float64(sum), float64(j-i+1)))
		m[str[j]] = j + 1
	}
	return sum
}
