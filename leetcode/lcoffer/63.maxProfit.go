package lcoffer

// min 记录最小的，然后和当日的比较，如果当日小于min，更新
// res 记录利润。用当日的值-min。算出新利润，如果利润大于res。更新
func maxProfit(prices []int) int {

	lens := len(prices)
	if lens <= 1 {
		return 0
	}
	res, min := 0, prices[0]
	for i := 1; i < lens; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		temp := prices[i] - min
		if temp > res {
			res = temp
		}
	}
	return res
}
