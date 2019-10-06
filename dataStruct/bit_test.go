package dataStruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBit(t *testing.T)  {
	a := assert.New(t)
	bit := NewBitMap(100)
	a.Equal(bit.Add(10),true)
	a.Equal(bit.Add(9),true)
	a.Equal(bit.Exist(10),true)
	a.Equal(bit.Exist(11),false)

}
