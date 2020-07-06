package math

func reverseBits(num uint32) uint32 {
	var res uint32
	var pow int = 31
	for num != 0 {
		// 把最后一位取出来，左移之后累加到结果中
		res += (num & 1) << pow
		num >>= 1
		pow--
	}
	return res
}
