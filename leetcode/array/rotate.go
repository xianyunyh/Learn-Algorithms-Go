package array

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]

func rotate2(nums []int, k int) {
	n := len(nums)
	//不移动
	if n <= 1 || (k%n) == 0 {
		return
	}
	temp := make([]int, len(nums))
	m := 0
	k = k % n
	for i := k; i < n; i++ {
		temp[i] = nums[m]
		m++
	}
	for i := 0; i < k; i++ {
		temp[i] = nums[n-k+i]
	}
	for k, v := range temp {
		nums[k] = v
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if n <= 1 || k == 0 {
		return
	}
	reserve(nums, 0, n-1)
	reserve(nums, 0, k-1)
	reserve(nums, k, n-1)
}

func reserve(nums []int, start, end int) {
	for ; start < end; start++ {
		nums[start], nums[end] = nums[end], nums[start]
		end--
	}
}
