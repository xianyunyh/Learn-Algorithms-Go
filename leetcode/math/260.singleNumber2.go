package math

func singleNumber2(nums []int) []int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	lastBit := res & (-res) // 最低位
	a, b := 0, 0
	for _, v := range nums {
		if v&lastBit == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}
