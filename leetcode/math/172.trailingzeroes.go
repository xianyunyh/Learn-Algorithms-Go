package math

func trailingZeroes(n int) int {
	res := 0
	divisor := 5
	for divisor <= n {
		res += n / divisor
		divisor *= 5
	}
	return res
}
