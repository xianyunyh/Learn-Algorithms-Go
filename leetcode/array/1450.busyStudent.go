package array

func busyStudent(startTime []int, endTime []int, queryTime int) int {

	res := 0
	for k, v := range startTime {
		end := endTime[k]
		if queryTime >= v && queryTime <= end {
			res++
		}

	}
	return res
}
