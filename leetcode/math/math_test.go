package math

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	reverse(1534236469)
}

func TestRomanToInt(t *testing.T) {
	a := assert.New(t)
	a.Equal(romanToInt("MMMXLV"), 3045)
	a.Equal(romanToInt("III"), 3)
	a.Equal(romanToInt("IV"), 4)
	a.Equal(romanToInt("LVIII"), 58)
	s := addBinary("11", "1")
	log.Println(s)

}

func TestMyAtoi(t *testing.T) {
	a := assert.New(t)
	a.Equal(myAtoi("-91283472332"), -2147483648)
}

func TestCcountPrimes(t *testing.T) {
	countPrimes(10)
}

func TestHammingWeight(t *testing.T) {
	hammingWeight(00000000000000000000000000001011)
}
