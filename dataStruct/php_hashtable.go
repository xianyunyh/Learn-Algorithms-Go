package dataStruct

type PHPHashTable struct {
	TableMask   uint64
	NumUsed     uint32
	NumElements uint32
	TableSize   uint32
	ArData      *Buckets
}

//
type PHPBucket struct {
	data    interface{}
	hashKey string
	next    *PHPBucket
	flag    uint8 // flag=1 被删除
}
type Buckets struct {
	Index []int
	Items []*PHPBucket
}

// 一个空表
func NewPHPHashTable() *PHPHashTable {
	idx := make([]int, 8)
	items := make([]*PHPBucket, 8)
	for i := 0; i < 8; i++ {
		idx[i] = -1
	}
	return &PHPHashTable{
		TableMask:   8,
		NumUsed:     0, //使用的个数 包含被删除的
		NumElements: 0, //真实的个数
		TableSize:   8, //总数
		ArData: &Buckets{
			Index: idx,
			Items: items,
		},
	}
}

//Insert 插入元素
func (h *PHPHashTable) Insert(key string, value interface{}) {

	//检查容量
	if h.NumUsed >= h.TableSize {
		h.Resize()
	}

	idx := h.HashCode(key)

	bucket := &PHPBucket{
		data:    value.(string),
		hashKey: key,
		next:    nil,
	}
	//冲突
	if h.ArData.Index[idx] != -1 {
		arrIdx := h.ArData.Index[idx]
		bucket.next = h.ArData.Items[arrIdx]
		h.ArData.Index[idx] = int(h.NumUsed)
		h.ArData.Items[h.NumUsed] = bucket
	} else {
		h.ArData.Index[idx] = int(h.NumUsed)
		h.ArData.Items[h.NumUsed] = bucket
	}
	h.NumUsed++
	h.NumElements++
}

//调整容量
func (h *PHPHashTable) Resize() {

	if h.NumUsed > h.NumElements+(h.NumElements>>5) {
		//删除里面已删除的元素，调整hash值
		for i := 0; i < int(h.TableSize); i++ {
			if h.ArData.Items[i].flag == 1 || h.ArData.Items[i] == nil {
				if int(h.TableSize) == i+1 {
					h.ArData.Items[i] = nil
					break
				}
				h.ArData.Items[i] = h.ArData.Items[i+1]
				h.ArData.Items[i+1] = nil
			}
		}
		h.Rehash()
		return
	}
	//扩容
	h.TableSize = h.TableSize * 2 //扩容两倍
	h.TableMask = uint64(h.TableSize)
	items := make([]*PHPBucket, h.TableSize)
	copy(items, h.ArData.Items)
	//items = append(items, h.ArData.Items[:]...)
	h.ArData.Index = make([]int, h.TableSize)
	h.ArData.Items = items
	//rehash 重建索引
	h.Rehash()

}

//再哈稀
func (h *PHPHashTable) Rehash() {

	//重置索引
	for i := range h.ArData.Index {
		h.ArData.Index[i] = -1
	}
	//遍历数组，重建索引
	for i, v := range h.ArData.Items {
		if v == nil {
			continue
		}
		t := h.HashCode(v.hashKey)

		//冲突
		if h.ArData.Index[t] != -1 {
			idx := h.ArData.Index[t]
			h.ArData.Items[i].next = h.ArData.Items[idx]
		} else {
			h.ArData.Index[t] = i
			h.ArData.Items[i].next = nil
		}

	}
}

// 计算索引
func (h *PHPHashTable) HashCode(key string) uint64 {
	return Time33(key) % (h.TableMask - 1)
}

//查找
func (h *PHPHashTable) Find(key string) interface{} {
	idx := h.HashCode(key)
	itemIdx := h.ArData.Index[idx]
	//没找到
	if itemIdx == -1 {
		return nil
	}
	//遍历链表
	item := h.ArData.Items[itemIdx]
	for item != nil {
		if item.hashKey == key {
			break
		}
		item = item.next
	}
	if item.flag != 1 {
		return item.data
	}
	return nil
}

// 删除的逻辑 就是设置flag 伪删除
func (h *PHPHashTable) Delete(key string) bool {
	hashNum := h.HashCode(key)
	if h.ArData.Index[hashNum] == -1 {
		return false
	}
	idx := h.ArData.Index[hashNum]
	item := h.ArData.Items[idx]
	for item != nil {
		if item.hashKey == key {
			item.flag = 1
			h.NumElements--
			return true
		}
		item = item.next
	}
	return false
}

//Count
func (h *PHPHashTable) Count() int {
	return int(h.NumElements)
}

//time 33 算法
func Time33(str string) uint64 {
	var hashMask uint64 = 5381
	for _, s := range str {
		c := uint64(s)
		hashMask += (c >> 5) + c // *33 = h×32+ h
	}
	return hashMask & 0x7FFFFFFF
}

func (h *PHPHashTable) Foreach(fn func(i int, val interface{})) {
	index := 0
	for _, v := range h.ArData.Items {
		if v == nil || v.flag != 0 {
			continue
		}
		fn(index, v.data)
		index++
	}
}
