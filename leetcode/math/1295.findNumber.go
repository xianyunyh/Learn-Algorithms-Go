package math

func findNumbers(nums []int) int {
	n := 0
	for _, v := range nums {
		if numLen(v)%2 == 0 {
			n = n + 1
		}
	}
	return n
}

func numLen(num int) int {
	i := 1
	for num >= 10 {
		num = num / 10
		i++
	}
	return i
}
