package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSort(t *testing.T) {
	//a := assert.New(t)
	arr := []int{1,2,3,5,3,7,9,4}
	res := CountSort(arr,1,9)
	a := assert.New(t)
	a.Equal(res,[]int{1,2,3,3,4,5,7,9})
}

func TestBubbleSort(t *testing.T) {
	arr := []int{2,1,5,8,3,6,9}
	a := assert.New(t)
	BubbleSort(arr)
	a.Equal(arr,[]int{1,2,3,5,6,8,9})
}

func TestInsertSort(t *testing.T) {
	arr := []int{2,1,5,8,3,6,9}
	InsertSort(arr)
	a := assert.New(t)
	a.Equal(arr,[]int{1,2,3,5,6,8,9})
}
func TestRadixSort(t *testing.T) {
	arr := []int{2,1,5,8,3,6,9}
	RadixSort(arr,1)
	t.Log(arr)
}

func TestBucketSort(t *testing.T) {
	arr := []int{2,1,5,8,3,6,9}
	res := BucketSort(arr)
	t.Log(res)
}

func TestShellSort(t *testing.T) {
	arr := []int{2,1,5,8,3,6,9}
	ShellSort(arr)
	t.Log(arr)
}