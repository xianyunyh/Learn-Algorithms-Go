package array

func createTargetArray(nums []int, index []int) []int {
	result := []int{}
	for k, v := range nums {
		i := index[k]
		left := result[:i]
		right := append([]int{}, result[i:]...)
		result = append(left, v)
		result = append(result, right[:]...)
	}
	return result
}
