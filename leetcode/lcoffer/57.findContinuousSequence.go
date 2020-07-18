package lcoffer

// 滑动窗口，当值大于目标，向左移动，小于向右移动，等于返回目标值
func findContinuousSequence(target int) [][]int {
	i, j := 1, 1
	sum := 0
	result := make([][]int, 0)
	for j <= target {
		if sum > target {
			sum = sum - i
			i++

		} else if sum < target {
			sum += j
			j++
		} else {

			temp := []int{}
			for n := i; n < j; n++ {
				temp = append(temp, n)
			}
			result = append(result, temp)
			sum = sum - i
			i++
		}

	}
	return result
}
