package hash

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	table := make(map[int]int)
	min := citations[0]
	max := citations[len(citations)-1]
	for max >= min {
		for _, v := range citations {
			if v-min >= 0 {
				if _, ok := table[min]; ok {
					table[min]++
				} else {
					table[min] = 1
				}
			}
		}
		min++
	}
	h := 1
	for k, v := range table {
		if v > h && k > 1 {
			h = v
		}
	}
	return h
}
