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
	table.Insert("aaaaatest", "22")
	table.Insert("c111ssdf", "33")
	table.Insert("100", "44")
	table.Insert("e", "55")
	table.Insert("2200", "66")
	table.Insert("g", "77")
	table.Insert("哈哈哈哈", "88")
	table.Insert("e", "eee")
	table.Insert("i", "88")
	a.Equal(table.Find("a"), "111")
	a.Equal(table.Find("哈哈哈哈"), "88")
	a.Equal(table.Count(), 9)
	table.Delete("a")
	a.Equal(table.Count(), 8)
	a.Equal(table.Find("a"), nil)
	table.Insert("a", "this is a")
	table.Foreach(func(k string, val interface{}) {
		log.Println(k, val)
	})
}
