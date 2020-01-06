package math

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		three := one + two
		one = two
		two = three
	}
	return two
}
