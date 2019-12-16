package dataStruct

import (
	"log"
	"testing"
)

func TestNewPHPHashTable(t *testing.T) {
	table := NewPHPHashTable()
	table.Insert("a","111")
	table.Insert("b","22")
	table.Insert("c","33")
	table.Insert("d","44")
	table.Insert("e","55")
	table.Insert("f","66")
	table.Insert("g","77")
	table.Insert("h","88")
	table.Insert("i","88")
	t.Log(table.Find("a"))
	t.Log(table.Find("b"))
	t.Log(table.Find("c"))
	t.Log(table.Find("d"))
	log.Println(table.Count())
	table.Delete("a")
	log.Println(table.Count())
	t.Log(table.Find("a"))
	table.Foreach(func(i int, val interface{}) {
		log.Println(i,val)
	})
}