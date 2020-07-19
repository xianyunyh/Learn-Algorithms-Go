package lcoffer

func singleNumber2(nums []int) int {
	one, two := 0, 0
	for _, v := range nums {
		one = (one ^ v) & (^two)
		two = (two ^ v) & (^one)
	}
	return one
}
