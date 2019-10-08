package sort

//计数排序
func CountSort(arr []int,min,max int) []int {
	lens := len(arr)
	counts := max - min + 1
	buffers := make([]int,counts)

	for _,v := range arr  {
		buffers[v - min] += 1
	}
	result := make([]int,lens)
	index := 0;
	for k,v := range buffers {
		for v > 0 {
			result[index] = k+min
			index++
			v--
		}
	}
	return result
}
