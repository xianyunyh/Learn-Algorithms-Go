package linklist

func distributeCandies(candies []int) int {
	counts := len(candies)
	setLen := 0
	set := make(map[int]struct{})
	for _, v := range candies {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			setLen += 1
		}
	}
	if counts/2 > setLen {
		return setLen
	}
	return counts / 2
}
