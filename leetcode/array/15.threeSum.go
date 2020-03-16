package array

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 || nums == nil {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}

		}
	}
	return result
}
