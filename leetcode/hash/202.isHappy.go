package hash

func sum(num int) int {
	sum := 0
	for num >= 10 {
		sum += (num % 10) * (num % 10)
		num = num / 10
	}
	sum += num * num
	return sum
}
func isHappy(n int) bool {
	m := make(map[int]struct{})
	for n != 1 {
		n = sum(n)
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
	}
	return true
}
