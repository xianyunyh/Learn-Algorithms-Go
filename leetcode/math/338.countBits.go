package math

func countBits(num int) []int {
	res := make([]int, num+1)

	for i := 0; i <= num; i++ {
		res[i] = count(i)
	}
	return res
}
func count(n int) (res int) {
	for n != 0 {
		n = n & (n - 1)
		res++
	}
	return
}
