package lcoffer

// 单调队列 题解参考
// https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/solution/mian-shi-ti-59-i-hua-dong-chuang-kou-de-zui-da-1-6/
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	for i := 0; i < k; i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[0 : len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res := []int{}
	res = append(res, queue[0])
	for i := k; i < len(nums); i++ {
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[0 : len(queue)-1]
		}
		queue = append(queue, nums[i])
		res = append(res, queue[0])
	}
	return res
}
