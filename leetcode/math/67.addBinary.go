package math

import "strconv"

func addBinary(a, b string) string {
	aBytes := []byte(a)
	bBytes := []byte(b)
	i := len(aBytes)
	j := len(bBytes)

	k := i
	if j > i {
		k = j
	}
	flag := 0
	result := make([]byte, k)
	for n := 1; n <= k; n++ {
		a1 := 0
		b1 := 0
		if i >= n {
			a1, _ = strconv.Atoi(string(aBytes[i-n]))
		}
		if j >= n {
			b1, _ = strconv.Atoi(string(bBytes[j-n]))
		}
		sum := a1 + b1 + flag
		p := sum % 2
		flag = sum / 2
		result[k-n] = byte(rune(p + 48))
		if n == k && flag == 1 {
			result = append([]byte{'1'}, result...)
		}
	}
	return string(result)
}
