package array

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	m1 := 2 << 31
	m2 := 2 << 31

	for _, v := range nums {
		if m1 >= v {
			m1 = v
		} else if m2 >= v {
			m2 = v
		} else {
			return true
		}
	}
	return false
}
