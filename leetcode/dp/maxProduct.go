package dp

//最大连续子数积
//递推公式
// dpMax[i] = max(arr[i],max(dpMax[i-1] * arr[i],dpMin[i-1] * arr[i]))
// dpMin[i] = min(arr[i],min(dpMax[i-1] * arr[i],dpMin[i-1] * arr[i]))
//max max(dpMax[i],dpMin[i]])
func maxProduct(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	maxArr := make([]int, n)
	minArr := make([]int, n)
	maxArr[0] = arr[0]
	minArr[0] = arr[0]
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	result := arr[0]
	for i := 1; i < n; i++ {
		maxTemp := arr[i] * maxArr[i-1]
		minTemp := arr[i] * minArr[i-1]
		maxArr[i] = max(arr[i], max(maxTemp, minTemp))
		minArr[i] = min(arr[i], min(maxTemp, minTemp))
		result = max(maxArr[i], minArr[i])
	}
	return result
}

func maxProduct2(nums []int) int {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}
