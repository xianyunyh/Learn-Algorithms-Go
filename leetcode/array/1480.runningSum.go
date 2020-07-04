package array

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	sum := 0
	for k, v := range nums {
		sum = v + sum
		res[k] = sum
	}
	return res
}
