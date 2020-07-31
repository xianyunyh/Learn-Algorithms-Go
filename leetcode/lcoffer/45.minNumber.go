package lcoffer

import (
	"sort"
	"strconv"
	"strings"
)

// x+y < y+x 102 <
// 排序判断规则： 设 nums任意两数字的字符串格式 x 和 y ，则
// 若拼接字符串 x + y > y + x则 m > n ；
// 反之，若 x + y < y + x 则 n < m ；
// 参考链接 https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/solution/mian-shi-ti-45-ba-shu-zu-pai-cheng-zui-xiao-de-s-4/
func minNumber(nums []int) string {
	stringArr := make([]string, len(nums))
	for i, v := range nums {
		stringArr[i] = strconv.Itoa(v)
	}
	sort.Slice(stringArr, func(i, j int) bool {
		return (stringArr[i] + stringArr[j]) < (stringArr[j] + stringArr[i])
	})
	return strings.Join(stringArr, "")
}
