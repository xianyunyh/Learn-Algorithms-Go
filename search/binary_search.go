package search

//二分查找 数组必须是有序的，先用中间的元素和目标进行比较 大于目标元素则往左查找，否则向右查找
func BinarySearch(arr []int ,target int ) bool  {
	min,max := 0,len(arr)-1
	result := false
	for min <= max  {
		mid := (max+min)/2
		if arr[mid] > target {
			max = mid-1
		} else if  arr[mid] < target {
			min = mid+1
		} else {
			result =  true
			break
		}
	}
	return result
}