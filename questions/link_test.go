package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xianyunyh/Learn-Algorithms-Go/dataStruct"
)

func TestFindReserveKNode(t *testing.T) {
	a := assert.New(t)
	l := dataStruct.NewLinkedList()
	data := []string{"1", "2", "3", "4", "5", "6"}
	for _, v := range data {
		l.Append(v)
	}
	ele := FindReserveKNode(l, 4)
	a.NotNil(ele)
	a.Equal(ele.Data, "3")
}

func TestReverseList(t *testing.T) {
	a := assert.New(t)
	l := dataStruct.NewLinkedList()
	data := []string{"1", "2", "3", "4", "5", "6"}
	for _, v := range data {
		l.Append(v)
	}
	ReverseList(l)
	a.Equal(l.Head.Next.Data, "5")
}
