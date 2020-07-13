package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v] = v
	}
	for _, v := range nums {
		a := target - v
		if _, ok := m[a]; ok {
			return []int{v, a}
		}
	}
	return nil
}
