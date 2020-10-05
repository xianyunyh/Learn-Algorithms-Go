package dataStruct

const (
	DELETE = 1
)

type PHPHashTable struct {
	tableMask   uint64
	numUsed     uint32
	numElements uint32
	tableSize   uint32
	arData      *arData
}

//
type PHPBucket struct {
	data    interface{}
	hashKey string
	next    *PHPBucket
	flag    uint8 // flag=1 被删除
}
type arData struct {
	indexs []int
	items  []*PHPBucket
}

// 一个空表
func NewPHPHashTable() *PHPHashTable {
	idx := make([]int, 8)
	items := make([]*PHPBucket, 8)
	for i := 0; i < 8; i++ {
		idx[i] = -1
	}
	//初始值为8 扩容是8*n
	return &PHPHashTable{
		tableMask:   8,
		numUsed:     0, //使用的个数 包含被删除的
		numElements: 0, //真实的个数
		tableSize:   8, //总数
		arData: &arData{
			indexs: idx,
			items:  items,
		},
	}
}

//Insert 插入元素
func (h *PHPHashTable) Insert(key string, value interface{}) {

	//检查容量
	if h.numUsed >= h.tableSize {
		h.reSize()
		h.rehash()
	}
	pos := h.findPos(key)
	if pos != -1 {
		item := h.arData.items[pos]
		for item != nil {
			if item.hashKey == key {
				item.data = value
				item.flag = 0
			}
			item = item.next
		}
		return
	}

	idx := h.HashCode(key)

	bucket := &PHPBucket{
		data:    value.(string),
		hashKey: key,
		next:    nil,
	}
	//冲突
	if h.arData.indexs[idx] != -1 {
		arrIdx := h.arData.indexs[idx]
		bucket.next = h.arData.items[arrIdx]
	}
	h.arData.indexs[idx] = int(h.numUsed)
	h.arData.items[h.numUsed] = bucket
	h.numUsed++
	h.numElements++
}

//调整容量
func (h *PHPHashTable) reSize() {

	if h.numUsed > h.numElements+(h.numElements>>5) {
		//删除里面已删除的元素，调整hash值
		for i := 0; i < int(h.tableSize); i++ {
			if h.arData.items[i].flag == DELETE || h.arData.items[i] == nil {
				if int(h.tableSize) == i+1 {
					h.arData.items[i] = nil
					break
				}
				h.arData.items[i] = h.arData.items[i+1]
				h.arData.items[i+1] = nil
			}
		}
		return
	}
	//扩容
	h.tableSize = h.tableSize * 2 //扩容两倍
	h.tableMask = uint64(h.tableSize)
	items := make([]*PHPBucket, h.tableSize)
	copy(items, h.arData.items)
	//items = append(items, h.arData.items[:]...)
	h.arData.indexs = make([]int, h.tableSize)
	h.arData.items = items
}

//再哈稀
func (h *PHPHashTable) rehash() {

	//重置索引
	for i := range h.arData.indexs {
		h.arData.indexs[i] = -1
	}
	//遍历数组，重建索引
	for i, v := range h.arData.items {
		if v == nil {
			continue
		}
		t := h.HashCode(v.hashKey)

		//冲突
		h.arData.items[i].next = nil
		if h.arData.indexs[t] != -1 {
			idx := h.arData.indexs[t]
			h.arData.items[i].next = h.arData.items[idx]
		}
		h.arData.indexs[t] = i
	}
}

// 计算索引
func (h *PHPHashTable) HashCode(key string) uint64 {
	return Time33(key) % (h.tableMask - 1)
}

func (h *PHPHashTable) findPos(key string) int {
	idx := h.HashCode(key)
	itemIdx := h.arData.indexs[idx]
	pos := -1
	//没找到
	if itemIdx == -1 {
		return pos
	}
	//遍历链表
	item := h.arData.items[itemIdx]
	for item != nil {
		if item.hashKey == key {
			break
		}
		item = item.next
	}
	if item != nil {
		return itemIdx
	}
	pos = -1
	return pos
}

//查找
func (h *PHPHashTable) Find(key string) interface{} {
	idx := h.HashCode(key)
	pos := h.arData.indexs[idx]
	//没找到
	if pos == -1 {
		return nil
	}
	//遍历链表
	item := h.arData.items[pos]
	for item != nil {
		if item.hashKey == key {
			break
		}
		item = item.next
	}
	if item != nil && item.flag != DELETE {
		return item.data
	}
	return nil
}

// 删除的逻辑 就是设置flag 伪删除
func (h *PHPHashTable) Delete(key string) bool {
	hashNum := h.HashCode(key)
	if h.arData.indexs[hashNum] == -1 {
		return false
	}
	pos := h.arData.indexs[hashNum]
	item := h.arData.items[pos]
	for item != nil {
		if item.hashKey == key {
			item.flag = 1
			h.numElements--
			return true
		}
		item = item.next
	}
	return false
}

//Count
func (h *PHPHashTable) Count() int {
	return int(h.numElements)
}

//time 33 算法
func Time33(str string) uint64 {
	var hashMask uint64 = 5381
	for _, s := range str {
		c := uint64(s)
		hashMask += (c << 5) + c // *33 = h×32+ h
	}
	return hashMask & 0x7FFFFFFF
}

func (h *PHPHashTable) Foreach(fn func(key string, val interface{})) {
	for _, v := range h.arData.items {
		if v == nil || v.flag == DELETE {
			continue
		}
		fn(v.hashKey, v.data)
	}
}
