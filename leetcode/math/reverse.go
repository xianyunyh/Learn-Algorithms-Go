package math

const INT_MAX = int32(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

//整数翻转
func reverse(x int) int32 {

	var rev int32 = 0

	for x != 0 {
		pop := x % 10 // 余数
		x /= 10
		if rev > INT_MAX/10 || (rev == INT_MAX/10 && pop > 7) {
			return 0
		}
		if rev < INT_MIN/10 || (rev == INT_MIN/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + int32(pop)
	}
	return rev
}
