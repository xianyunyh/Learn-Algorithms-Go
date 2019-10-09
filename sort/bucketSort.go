package sort

// 桶排序
func BucketSort(arr []int) ([]int)  {
	min,max := arr[0],arr[0]
	for _,v := range arr {
		if v > max {
			max = v
		}else if v < min {
			min = v
		}
	}
	//buckets
	lens := max - min + 1
	bucketNum := 5
	buckets := make([][]int,bucketNum)
	step := (lens / bucketNum) + 1

	for _,v := range arr {
		index := (v - min) / step
		buckets[index] = append(buckets[index],v)
	}
	resultArr := make([]int,0)
	for _,bucket := range buckets  {
		resultArr = append(resultArr,BubbleSort(bucket)...)
	}
	return resultArr
}
