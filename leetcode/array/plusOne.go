package array

//PlusOne  给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。
//最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
//你可以假设除了整数 0 之外，这个整数不会以零开头
// 关键信息 1、 非负整数数据 2、最高位放首位，从高位到低的数组 3、数组只存一个数据 0-9 4、结果不能以0开头
//可能出现的情况
// 不进位 如 20
// 只末尾进位，如19
// 末尾+中间连续进位，如199
// 末尾+中间+首位进位，如99
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		//不是9 不进位
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	//全为9
	newArr := make([]int, 0)
	newArr = append(newArr, 1)
	newArr = append(newArr, digits...)
	return newArr
}
