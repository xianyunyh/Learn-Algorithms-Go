package questions

//单链表结构
type Node struct {
	Val  int
	Next *Node
}

//使用单链表实现lru缓存
//1. 如果数据存在链表中，那么遍历整个链表删除节点，再插入到头部
//2. 如果没存在链表
// 2.1 如果已满，淘汰链表尾部的节点，插入到头部
// 2.2 如果未满，直接插入链表头
type LRUCache struct {
	Head     *Node
	Capacity int
	Lens     int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Lens:     0,
		Capacity: capacity,
		Head:     nil,
	}
}

func (l *LRUCache) Put(val int) {
	if l.Lens == 0 {
		l.insertHead(val)
		return
	}
	prevNode := l.findPrevNode(val)
	if prevNode != nil {
		l.deleteNode(prevNode)
		l.insertHead(val)
	} else {
		//需要淘汰
		if l.Full() {
			l.deleteLast()
		}
		l.insertHead(val)
	}
}

func (l *LRUCache) insertHead(val int) {
	node := &Node{
		Val: val,
	}
	head := l.Head
	node.Next = head
	l.Head = node
	l.Lens++
}

func (l *LRUCache) findPrevNode(val int) *Node {
	if l.Lens == 0 {
		return nil
	}
	prev := l.Head
	current := prev.Next
	for current != nil {
		if current.Val == val {
			return prev
		}
		prev = current
		current = current.Next
	}
	return nil
}
func (l *LRUCache) existNode(val int) bool {
	node := l.Head
	for node != nil {
		if node.Val == val {
			return true
		}
		node = node.Next
	}
	return false
}

func (l *LRUCache) deleteNode(prev *Node) {
	if prev == nil || prev.Next == nil {
		return
	}
	prev.Next = nil
}
func (l *LRUCache) deleteLast() {
	if l.Head == nil {
		return
	}
	prev := l.Head
	current := prev.Next
	for current != nil && current.Next != nil {
		prev = current
		current = current.Next
	}
	prev.Next = nil
	l.Lens--
}

func (l *LRUCache) Full() bool {
	if l.Lens > 0 && l.Lens >= l.Capacity {
		return true
	}
	return false
}
