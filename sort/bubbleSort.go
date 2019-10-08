package sort

func BubbleSort(arr []int) {
	lens := len(arr)
	//i := 0
	for i := 0; i < lens; i++ {
		for j := 0; j < lens-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
}
