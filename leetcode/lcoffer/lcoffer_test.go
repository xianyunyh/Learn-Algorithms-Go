package lcoffer

import "testing"

func TestSpiralOrder(t *testing.T) {
	spiralOrder([][]int{
		{2, 5, 8},
		{4, 0, -1},
	})
}

func TestSlice(t *testing.T) {
	str := "hello world"
	t.Log(str[6:12])
}
