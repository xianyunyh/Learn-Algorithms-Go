package math

// N 和 (N-1) 做与操作 可以将最后一位1置零
func HammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}
