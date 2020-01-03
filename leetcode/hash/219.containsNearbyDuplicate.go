package hash

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]struct{})

	for i, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
		if len(set) > k {
			delete(set, nums[i-k])
		}
	}
	return false
}
