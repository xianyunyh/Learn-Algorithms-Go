package linklist

//实现一个LRU（最近最少使用淘汰）使用hashTable + Double link

type Node struct {
	key   int
	value int
	Prev  *Node
	Next  *Node
}

type DoublueLink struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLink() *DoublueLink {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	link := &DoublueLink{
		head: head,
		tail: tail,
	}
	return link
}

func (self *DoublueLink) createNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
		Prev:  nil,
		Next:  nil,
	}
}
func (self *DoublueLink) AddFirst(node *Node) *Node {
	head := self.head
	node.Prev = head
	node.Next = head.Next
	head.Next.Prev = node
	head.Next = node
	self.size++
	return node
}

func (self *DoublueLink) Remove(node *Node) {
	if self.size == 0 {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	self.size--
}

func (self *DoublueLink) RemoveLast() *Node {
	if self.size == 0 {
		return nil
	}
	node := self.tail.Prev
	self.Remove(node)
	return node
}
func (self *DoublueLink) Total() int {
	return self.size
}

type LRUCache struct {
	capacity int
	store    map[int]*Node
	link     *DoublueLink
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		store:    make(map[int]*Node),
		capacity: capacity,
		link:     NewDoubleLink(),
	}
}

func (this *LRUCache) Get(key int) int {
	store := this.store
	if _, ok := store[key]; ok {
		node := store[key]
		this.link.Remove(node)
		node = this.link.AddFirst(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	store := this.store
	//已经存在
	node := &Node{key: key, value: value}
	if _, ok := store[key]; ok {
		this.link.Remove(store[key])
		node = this.link.AddFirst(node)
		this.store[key] = node
		return
	}
	//已满
	if this.capacity == this.link.Total() {
		last := this.link.RemoveLast()
		delete(this.store, last.key)
		node = this.link.AddFirst(node)
		this.store[key] = node
		return
	}
	this.link.AddFirst(node)
	this.store[key] = node
}
