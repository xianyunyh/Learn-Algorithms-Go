package lcoffer

import (
	"sort"
)

// 解题思路 先排序，王牌肯定在最前面。遍历，算王牌的数量。
// 如果中间有两个相同的牌，则不是顺子 0 0 1 1 2 3
// 因为有五张牌，最大的跨度就是5，如果王牌之后的最小牌和第五张牌大于5 则不是顺子 0 0 1 2 6不是顺子  0 0 1 2 5 是顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)
	qLen := 0
	for i := 0; i < 4; i++ {
		// 算王牌
		if nums[i] == 0 {
			qLen++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}

	return nums[4]-nums[qLen] < 5

}
