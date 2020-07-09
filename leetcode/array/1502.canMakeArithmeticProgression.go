package array

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	lens := len(arr)
	if lens <= 2 {
		return true
	}
	sort.Ints(arr)
	x := arr[1] - arr[0]

	for i := 1; i < lens-1; i++ {
		if arr[i+1]-arr[i] != x {
			return false
		}
	}
	return true

}

// 利用等差公式 s = (a1 +an) *n /2
func canMakeArithmeticProgression2(arr []int) bool {
	lens := len(arr)
	min, max := arr[0], arr[0]
	sum := arr[0]
	for i := 1; i < lens; i++ {
		if arr[i] <= min {
			min = arr[i]
		}
		if arr[i] >= max {
			max = arr[i]
		}
		sum += arr[i]
	}
	return ((min + max) * lens) == 2*sum

}
