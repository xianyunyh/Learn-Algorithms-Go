package string

func reverseLeftWords(s string, n int) string {
	if len(s) <= 1 {
		return s
	}
	sLen := len(s)
	n = n % sLen
	if n == 0 {
		return s
	}
	result := make([]byte, sLen)

	for i := 0; i < n; i++ {
		result[sLen-n+i] = s[i]
	}
	for i := 0; i < sLen-n; i++ {
		result[i] = s[n+i]
	}
	return string(result)
}
