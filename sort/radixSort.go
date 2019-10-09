package sort

import (
	"math"
)

func RadixSort(arr []int,k int) {
	for i := 1; i <= k;i++{
		buckets := make([][]int,10)
		for j,v := range arr {
			//计算index
			index := arr[j]/int(math.Pow10(i-1)) - (arr[j]/int(math.Pow10(i)))*10
			buckets[index] = append(buckets[index],v)
		}
		arrIdx := 0
		for _,v1 := range buckets {
			for _,item := range v1  {
				arr[arrIdx] = item
				arrIdx++
			}
		}
	}
}
