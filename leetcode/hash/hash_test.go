package hash

import (
	"testing"

	"log"

	"github.com/stretchr/testify/assert"
)

func TestNumJewelsInStones(t *testing.T) {
	a := assert.New(t)

	a.Equal(NumJewelsInStones("aA", "aAAbbbb"), 3)
	a.Equal(NumJewelsInStones("z", "ZZ"), 0)
	a.Equal(countSegments("Hello, my name is John"), 5)
	log.Println(twoSum([]int{3, 2, 4}, 6))
}

func TestIsValidSudoku(t *testing.T) {
	bytes1 := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	bytes2 := [][]byte{
		[]byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	a := assert.New(t)
	a.Equal(isValidSudoku(bytes1), true)
	a.Equal(isValidSudoku(bytes2), false)
}
