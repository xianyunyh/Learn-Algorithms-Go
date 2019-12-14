package hash

import "testing"
import "github.com/stretchr/testify/assert"

func TestNumJewelsInStones(t *testing.T) {
	a := assert.New(t)

	a.Equal(NumJewelsInStones("aA","aAAbbbb"),3)
	a.Equal(NumJewelsInStones("z","ZZ"),0)
	a.Equal(countSegments("Hello, my name is John"),5)
}
