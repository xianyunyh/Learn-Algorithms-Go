package sort

//基本思路是：
// 将待排序元素划分到不同的痛。先扫描一遍序列求出最大值 maxV 和最小值 minV ，设桶的个数为 k ，
// 则把区间 [minV, maxV] 均匀划分成 k 个区间，每个区间就是一个桶。将序列中的元素分配到各自的桶。
// 对每个桶内的元素进行排序。可以选择任意一种排序算法。
// 将各个桶中的元素合并成一个大的有序序列。
// 假设数据是均匀分布的，则每个桶的元素平均个数为 n/k 。
// 假设选择用快速排序对每个桶内的元素进行排序，那么每次排序的时间复杂度为 O(n/klog(n/k)) 。
// 总的时间复杂度为 O(n)+O(m)O(n/klog(n/k)) = O(n+nlog(n/k)) = O(n+nlogn-nlogk 。
// 当 k 接近于 n 时，桶排序的时间复杂度就可以金斯认为是 O(n) 的。即桶越多，时间效率就越高，而桶越多，空间就越大
func BubbleSort(arr []int) ( []int) {
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
	return arr
}
