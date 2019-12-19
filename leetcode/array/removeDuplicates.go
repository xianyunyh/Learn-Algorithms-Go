package array

import "sort"

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

//利用快慢指针的方法
//
func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow + 1
}

func removeDuplicates2(nums []int) int {
	tempMap := make(map[int]int, len(nums))
	for k, v := range nums {
		//已经存在
		if _, ok := tempMap[v]; ok {
			continue
		}
		tempMap[v] = k
	}
	keys := []int{}
	for k, _ := range tempMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, v := range keys {
		nums[i] = v
	}
	return len(keys)
}
