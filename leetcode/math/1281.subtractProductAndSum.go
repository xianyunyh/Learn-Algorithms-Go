package math

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n >= 10 {
		m := n % 10
		n = n / 10
		sum += m
		product = m * product
	}
	sum += n
	product = product * n
	return product - sum
}
