package hash

import (
	"strconv"
)

// 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
//
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
//解题思路
//利用hash表 记录每行每列和 每个3*3宫格 每个数字出现的数字。从头遍历
func isValidSudoku(board [][]byte) bool {
	var row = [9]map[int]int{}
	var column = [9]map[int]int{}
	var box = [9]map[int]int{}

	for i := 0; i < 9; i++ {
		row[i] = make(map[int]int)
		column[i] = make(map[int]int)
		box[i] = make(map[int]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			char := board[i][j]
			if char != '.' {
				num, _ := strconv.Atoi(string(char))
				// 计算box所在的索引
				boxIdx := (i/3)*3 + j/3
				if _, ok := row[i][num]; ok {
					row[i][num]++
				} else {
					row[i][num] = 1
				}
				if _, ok := column[j][num]; ok {
					column[j][num]++
				} else {
					column[j][num] = 1
				}
				if _, ok := box[boxIdx][num]; ok {
					box[boxIdx][num]++
				} else {
					box[boxIdx][num] = 1
				}
				if row[i][num] > 1 || column[j][num] > 1 || box[boxIdx][num] > 1 {
					return false
				}
			}
		}
	}
	return true
}
