package lcoffer

//解题思路：每一行是最后一列是最大的，所以从左上角开始扫描
// 如果该值大于target 跳过一行，
// 如果该值小于target 跳过一列
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, colums := len(matrix)-1, len(matrix[0])-1
	row, colum := 0, colums
	for row <= rows && colum >= 0 {
		if matrix[row][colum] == target {
			return true
		}
		if matrix[row][colum] > target {
			colum--
		} else {
			row++
		}
	}
	return false
}
