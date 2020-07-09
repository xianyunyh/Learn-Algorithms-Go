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
