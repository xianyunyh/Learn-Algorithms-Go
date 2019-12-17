package dataStruct

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPHPHashTable(t *testing.T) {
	a := assert.New(t)
	table := NewPHPHashTable()
	table.Insert("a", "111")
	table.Insert("b", "22")
	table.Insert("c", "33")
	table.Insert("d", "44")
	table.Insert("e", "55")
	table.Insert("f", "66")
	table.Insert("g", "77")
	table.Insert("h", "88")
	table.Insert("i", "88")
	a.Equal(table.Find("a"), "111")
	a.Equal(table.Find("b"), "22")
	a.Equal(table.Count(), 9)
	table.Delete("a")
	a.Equal(table.Count(), 8)
	a.Equal(table.Find("a"), nil)
	table.Foreach(func(i int, val interface{}) {
		log.Println(i, val)
	})
}
