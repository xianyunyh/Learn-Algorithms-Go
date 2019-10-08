package sort

//插入排序
func InsertSort(arr []int)  {
	lens := len(arr)
	var i,j int
	for i = 1;i < lens ; i++  {
		temp := arr[i]
		for j = i -1 ; j >= 0 && temp < arr[j] ; j--  {
			arr[j+1] =  arr[j]
		}
		arr[j+1] = temp
	}
}
