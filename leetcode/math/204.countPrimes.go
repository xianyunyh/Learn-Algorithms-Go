package math

// 埃拉托斯特尼筛法

func countPrimes(n int) int {
	nums := make([]bool, n)
	for i := 2; i < n; i++ {
		if nums[i] == false {
			for j := 2; j*i < n; j++ {
				nums[j*i] = true
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if nums[i] == false {
			count++
		}
	}
	return count
}

//最简单的 效率最差
func countPrimes2(n int) int {
	isPrime := func(n int) bool {
		for j := 2; j < n; j++ {
			if n%j == 0 {
				return false
			}
		}
		return true
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}
