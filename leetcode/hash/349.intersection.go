package hash

// 解题思路，利用两个set
func intersection(nums1 []int, nums2 []int) []int {

	table1 := make(map[int]struct{})
	table2 := make(map[int]struct{})

	for _, v := range nums1 {
		if _, ok := table1[v]; !ok {
			table1[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := table1[v]; ok {
			if _, ok2 := table2[v]; !ok2 {
				table2[v] = struct{}{}
			}
		}
	}
	var reuslt []int

	for k, _ := range table2 {
		reuslt = append(reuslt, k)
	}
	return reuslt
}
