package algorithms

// Time33是字符串哈希函数，现在几乎所有流行的HashMap都采用了DJB Hash Function，
// 俗称“Times33”算法。Times33的算法很简单，就是不断的乘33。
func Time33(str string) uint64  {
	var  hashMask uint64  = 5381
	for _,s := range str {
		c := uint64(s)
		hashMask += ( c >> 5) + c // *33 = h×32+ h
	}
	return hashMask & 0x7FFFFFFF
}