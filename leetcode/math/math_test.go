package math

import "testing"

import "github.com/stretchr/testify/assert"

func TestReverse(t *testing.T) {
	reverse(1534236469)
}

func TestRomanToInt(t *testing.T) {
	a := assert.New(t)
	a.Equal(romanToInt("MMMXLV"), 3045)
	a.Equal(romanToInt("III"), 3)
	a.Equal(romanToInt("IV"), 4)
	a.Equal(romanToInt("LVIII"), 58)

}
