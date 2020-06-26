package hash

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素
//用hash来判断，注意的是需要判断两个index 不相同
func twoSum(nums []int, target int) []int {

	numsMap := make(map[int]int)

	for i, v := range nums {
		numsMap[v] = i
	}
	for i, v := range nums {
		n := target - v
		if j, ok := numsMap[n]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
