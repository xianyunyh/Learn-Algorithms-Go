package math

// 解题思路 1个数是4的幂 一定是2的幂，先判断是不是2的幂，4的幂的二进制1落在奇数位上。
// 0x55555555 010101010101
func isPowerOfFour(num int) bool {
	return num > 0 && (num&(num-1)) == 0 && (num&0x55555555) == num
}
