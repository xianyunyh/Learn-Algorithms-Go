package questions

import (
	"github.com/xianyunyh/Learn-Algorithms-Go/dataStruct"
)

//输入一个单链表，输出链表的倒数第K个节点

func FindReserveKNode(l *dataStruct.LinkedList, k int) *dataStruct.Node  {
	if l.Head == nil || k == 0 {
		return nil
	}
	slowNode := l.Head
	fastNode := l.Head
	for i := 1; i <= k ; i++  {
		if  fastNode != nil {
			fastNode = fastNode.Next
		}
	}
	for fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
	return slowNode
}

// 判断链表是否有环
/**
快慢指针 快的先走一步 然后慢的再走 如果有环 两个指针 一定在环内相遇 没环在终点相遇
 */

func LinkExistLoop(l *dataStruct.LinkedList) bool  {
	fast,slow := l.Head,l.Head
	for slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//节点相遇
		if slow == fast {
			return true
		}
	}
	return false
}

func getMeetingNode(l *dataStruct.LinkedList) *dataStruct.Node  {
	fast,slow := l.Head,l.Head
	for slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//节点相遇
		if slow == fast {
			return slow
		}
	}
	return nil
}

// 找到环的入口
func GetEntryNodeOfLoop(l *dataStruct.LinkedList) *dataStruct.Node  {
	meetingNode := getMeetingNode(l)
	if meetingNode == nil {
		return  nil
	}
	first,sencond := meetingNode,l.Head
	for first != sencond {
		first = first.Next
		sencond = sencond.Next
	}
	return sencond
}

func GetLoopLen(l *dataStruct.LinkedList) int  {
	if l.Head == nil {
		return 0
	}
	slow,fast := l.Head,l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	slow = slow.Next
	fast = fast.Next.Next
	lens := 1
	for fast != slow {
		slow = slow.Next
		fast = fast.Next.Next
		lens ++
	}
	return  lens
}
//翻转单链表 三个指针 pre current next
// 第一次让prev = nil 然后每次调整current的指向
func ReverseList(l *dataStruct.LinkedList)  {
	current := l.Head
	var prev *dataStruct.Node = nil
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current =	next
	}
	l.Head = prev
}