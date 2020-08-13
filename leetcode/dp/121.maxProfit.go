package dp

// 股票最佳卖出时机
// dp[i] = max((dp[i+1] -dp[i]) + dp[i-1],0)
// maxdp[i] = max(maxdp[i-1],dp[i])
func maxProfit(nums []int) int {
	profit, last := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		last = max(0, diff+last)
		profit = max(profit, last)
	}
	return profit
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
