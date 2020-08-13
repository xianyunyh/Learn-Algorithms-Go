package dp

// 53. 最大子序和
// dp[i] = max(dp[i-1]+nums[i],nums[i])
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//加
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	// 记录最大和
	maxSum := nums[0]
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		// 1,2-5 100
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}
