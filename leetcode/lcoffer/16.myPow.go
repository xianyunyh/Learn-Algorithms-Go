package lcoffer

// 一个数的n次幂 等于 a *...an
// a^5 = a^1+a^2+a^2
func myPow(x float64, n int) float64 {
	if x == 0 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var res float64 = 1
	for n > 0 {
		if n&1 > 0 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res

}
