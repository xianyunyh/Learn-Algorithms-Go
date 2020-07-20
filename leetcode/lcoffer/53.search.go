package lcoffer

// 解题思路 寻找target最后一次出现的位置和target-1最后一次的位置，中间差就是target出现的次数
func search(nums []int, target int) int {
	return find(nums, target) - find(nums, target-1)
}

func find(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}
