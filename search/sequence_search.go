package search

//顺序查找  重头到尾依次查找。 时间复杂度为O（n）
func SequenceSearch(arr []int, target int) bool  {
	result := false
	for _,v := range arr {
		if v == target {
			result = true
			break
		}
	}
	return result
}
