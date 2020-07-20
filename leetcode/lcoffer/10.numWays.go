package lcoffer

// 本质是个fib数列
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	one := 1
	two := 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = (one + two) % 1000000007
		one = two
		two = sum
	}
	return sum % 1000000007
}
