package lcoffer

// 解题思路，计算 result[i] =  a[i] 左边的元素 * a[i] 右边的元素
func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	result := make([]int, len(a))
	left := make([]int, len(a))
	right := make([]int, len(a))
	left[0] = 1
	for i := 1; i < len(a); i++ {
		left[i] = left[i-1] * a[i-1]
	}
	right[len(a)-1] = 1
	for i := len(a) - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}
	for i := 0; i < len(a); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

// 暴力法
func constructArr2(a []int) []int {
	result := make([]int, len(a))

	for k, _ := range a {
		temp := 1
		for k1, v1 := range a {
			if k1 == k {
				continue
			}
			temp *= v1
		}
		result[k] = temp
	}
	return result
}
