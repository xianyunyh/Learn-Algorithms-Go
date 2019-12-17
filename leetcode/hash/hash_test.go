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
