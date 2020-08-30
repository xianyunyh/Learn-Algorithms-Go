package lcoffer

import "sort"

//利用hash表
// 排序，超过一半的数，一定在中间的位置
func majorityElement(nums []int) int {
	table := make(map[int]int)

	for _, v := range nums {
		table[v] += 1
	}
	for k, v := range table {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

// O(N*Log(N))
func majorityElement2(nums []int) int {

	sort.Ints(nums)
	return nums[len(nums)/2]
}
