package stack

import "testing"

func TestIsValid(t *testing.T) {
	isValid("})")
}

func TestCQueue(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(5)
	obj.AppendTail(2)
	a := obj.DeleteHead()
	t.Log(a)
	a = obj.DeleteHead()
	t.Log(a)
}
