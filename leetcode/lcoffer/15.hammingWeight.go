package lcoffer

func hammingWeight(num uint32) int {
	result := 0

	for num > 0 {
		result += int(num & 1)
		num >>= 1
	}
	return result
}

// num  & num -1 消去最后一个1
func hammingWeight2(num uint32) int {
	result := 0
	for num > 0 {
		result++
		num &= num - 1
	}
	return result
}
