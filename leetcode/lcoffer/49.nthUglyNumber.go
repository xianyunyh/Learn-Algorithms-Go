package lcoffer

// 解题思路
// 声明三个指针 分别记录 2 ，3，5的集合，丑数一定是2，3，5中间的数相乘。那么一定会有以下集合
// [1*2,2*2,3*2,4*2,5*2...] [1*3....5*3...] [1*5...5*5]
// 将三个数组进行合并，就是最终的结果，由于会出现重复，比如2*3 3*2,则需要去重
func nthUglyNumber(n int) int {
	dp := []int{}
	dp = append(dp, 1)
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		temp := min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		if temp == dp[p2]*2 {
			p2++
		}
		if temp == dp[p3]*3 {
			p3++
		}
		if temp == dp[p5]*5 {
			p5++
		}
		dp = append(dp, temp)
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
