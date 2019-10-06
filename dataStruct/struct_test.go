package dataStruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	a := assert.New(t)
	list := NewLinkedList()
	list.Append("1")
	list.Append("2")
	a.Equal(list.lens,2)
	a.Equal(list.Search("1"),true)
	a.Equal(list.Search("2"),true)
	a.Equal(list.Search("3"),false)
	a.Equal(list.Remove("2"),true)
	a.Equal(list.Search("1"),true)
	a.Equal(list.Lens(),1)
}

func TestQueue(t *testing.T)  {
	a := assert.New(t)
	q := NewQueue(5)
	items := []string{"hello","world","china","php","study","word"}
	for _,v := range items  {
		q.EnQueue(v)
	}
	a.Equal(q.Lens(),5)
	a.Equal(q.DeQueue(),"hello")
}
func TestStack(t *testing.T)  {
	a := assert.New(t)
	s := NewStack(5)
	items := []string{"hello","world","china","php","study","word"}
	for _,v := range items  {
		s.Push(v)
	}
	a.Equal(s.Pop(),"study")
	a.Equal(s.Pop(),"php")
	a.Equal(s.Pop(),"china")
}
func TestHashTable(t *testing.T)  {
	a := assert.New(t)
	hash := NewHashTable(1024)
	hash.Insert("wo","1")
	//碰撞
	hash.Insert("ow","22")
	hash.Insert("wod","2")
	hash.Insert("wod","1122")
	//t.Log(hash.Get("wod"))
	hash.Insert("中国","1244")
	a.Equal(hash.Get("wo"),"1")
	a.Equal(hash.Get("wod"),"1122")
}

func TestSet(t *testing.T)  {
	a := assert.New(t)
	set := NewSet()
	a.Equal(set.Add("111"),true)
	a.Equal(set.Add("2222"),true)
	a.Equal(set.Add("111"),false)
	a.Equal(set.Lens(),2)
}