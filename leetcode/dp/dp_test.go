package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	num := maxProduct([]int{1, -2, 3, -4, -2})
	a := assert.New(t)
	a.Equal(num, 24)
}

func TestMaxSubArr(t *testing.T) {
	num := maxSubArray([]int{1, -2, 5, 10, -21})
	assert.Equal(t, num, 15)
}

func TestMaxProfit(t *testing.T) {
	num := maxProfit([]int{7, 1, 5, 3, 6, 4})
	num2 := maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, num, 5)
	assert.Equal(t, num2, 0)
}
