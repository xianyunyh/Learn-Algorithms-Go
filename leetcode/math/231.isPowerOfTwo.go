package math

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	r := 0
	for n > 0 {
		r += (n & 1)
		n >>= 1
	}
	return r == 1
}
