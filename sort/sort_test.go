package sort

import (
	"testing"
)

func TestCountSort(t *testing.T) {
	//a := assert.New(t)
	arr := []int{1,2,3,5,3,7,9,4}
	res := CountSort(arr,1,9)
	for _,v := range res {
		t.Log(v)
	}
}
