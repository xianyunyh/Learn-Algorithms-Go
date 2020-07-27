package lcoffer

// 思路 对于一个旋转数组，前面一段肯定是递增，后面是递减，从头开始遍历，然后遇到nums[i+1] < nums[i] i+1 就是递减的开始位置
func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return numbers[i]
		}
	}
	return numbers[0]
}
