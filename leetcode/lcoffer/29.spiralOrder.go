package lcoffer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	k := 0
	res := make([]int, (bottom+1)*(right+1))

	for left <= right && top <= bottom {
		//上边一行
		for i := left; i <= right; i++ {
			res[k] = matrix[top][i]
			k++
		}
		top++

		//右边一列
		for i := top; i <= bottom; i++ {
			res[k] = matrix[i][right]
			k++
		}
		right--

		//下边一行
		for i := right; i >= left && top <= bottom; i-- {
			res[k] = matrix[bottom][i]
			k++
		}
		bottom--

		//左边一列
		for i := bottom; i >= top && left <= right; i-- {
			res[k] = matrix[i][left]
			k++
		}
		left++
	}
	return res
}
