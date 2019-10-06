package search

import "log"

//   插入查找：选择下标 = left + (right - left) * (key - arr[left])/(arr[right] - arr[left])
//    插值查找是根据要查找的关键字key与查找表中最大最小记录的关键字比较后的
//    查找方法，其核心就在于插值的计算公式 (key-a[low])/(a[high]-a[low])*(high-low)。
//   复杂度分析
//   时间复杂性：如果元素均匀分布，则O（log log n）），在最坏的情况下可能需要 O（n）。
//   空间复杂度：O（1）。
func DvalueSearch(arr []int,target int ) bool  {
	min,max := 0,len(arr)-1
	i := 0;
	result := false;
	for min <= max  {
		i++
		mid := min + (max-min)*((target-arr[min])/(arr[max]-arr[min]))
		if arr[mid] > target {
			max = mid -1
		} else if arr[mid] < target {
			min = mid + 1
		} else {
			result = true
			break
		}
	}
	log.Println(i)
	return result
}
