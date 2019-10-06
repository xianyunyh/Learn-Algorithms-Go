package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7}
	a := assert.New(t)
	a.Equal(BinarySearch(arr,5),true,"查找出错")
	a.Equal(BinarySearch(arr,10),false,"查找出错")
	a.Equal(BinarySearch(arr,-1),false,"查找出错")
}

func TestSequenceSearch(t *testing.T) {
	arr := []int{1,2,3,4,6,7,10}
	a := assert.New(t)
	a.Equal(SequenceSearch(arr,1),true)
	a.Equal(SequenceSearch(arr,8),false)
}

func TestDvalueSearch(t *testing.T) {
	arr := []int{1,2,3,4,6,7,10}
	a := assert.New(t)
	a.Equal(DvalueSearch(arr,2),true)
	a.Equal(DvalueSearch(arr,11),false)
}