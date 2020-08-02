package lcoffer

// 题解 https://leetcode-cn.com/problems/jian-sheng-zi-lcof/solution/mian-shi-ti-14-i-jian-sheng-zi-tan-xin-si-xiang-by/
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	pow := func(a, b int) int {
		res := 1
		for b > 0 {
			res *= a
			b--
		}
		return res
	}
	if b == 0 {
		return pow(3, a)
	}
	if b == 1 {
		return pow(3, a-1) * 4
	}
	return pow(3, a) * 2

}

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	pow := func(a, b int) int {
		res := 1
		for b > 0 {
			res = (a * res) % 1000000007
			b--
		}
		return res
	}
	if b == 0 {
		return pow(3, a) % 1000000007
	}
	if b == 1 {
		return (pow(3, a-1) * 4) % 1000000007
	}
	return (pow(3, a) * 2) % 1000000007

}
