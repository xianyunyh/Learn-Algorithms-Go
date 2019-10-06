package dataStruct

type HashTable struct {
	lens    int
	mask    int
	buckets []*Bucket
}
type Bucket struct {
	data    string
	hashKey string
	idx     int
	next    *Bucket
}

//创建一个哈希表
func NewHashTable(size int) *HashTable {
	return &HashTable{
		lens:    0,
		mask:    size,
		buckets: make([]*Bucket, size),
	}
}
//todo 很简单的hash算法
func (h *HashTable) hashFun(s string) int {
	//简单的哈希算法
	sum := rune(0);
	for i := 0;i < len(s);i++ {
		sum += rune(s[i])
	}
	k := int(sum)
	if k > h.mask {
		return  k % h.mask
	}
	return h.mask % k
}

// 插件值
func (h *HashTable) Insert(key, value string) {
	bucket := &Bucket{
		data:    value,
		hashKey: key,
	}
	idx := h.hashFun(key)
	if len(h.buckets) < idx || h.buckets[idx] == nil {
		h.buckets[idx] = bucket
		bucket.next = nil
	} else {
		//拉链 解决冲突
		head := h.buckets[idx]
		for head.next != nil {
			if head.hashKey != key {
				head = head.next
			}
		}
		//重复键名
		if head.hashKey == key {
			head.data = value
			return
		}
		head.next = bucket
	}
	h.lens++
}

//根据key获取值
func (h *HashTable) Get(key string) string {
	idx := h.hashFun(key)
	if h.buckets[idx] == nil {
		return ""
	}
	head := h.buckets[idx]
	for head.hashKey != key {
		head = head.next
	}
	return head.data
}

func (h *HashTable) Del(key string) bool {
	idx := h.hashFun(key)
	if h.buckets[idx] == nil {
		return false
	}
	node := h.buckets[idx]
	for node.hashKey != key {
		node = node.next
	}
	return false
}

//遍历哈希表
func (h *HashTable) Scan(fn func(k, v string)) {
	for i := 0; i < h.mask; i++ {
		if h.buckets[i] == nil {
			continue
		}
		bucket := h.buckets[i]
		for bucket.next != nil {
			fn(bucket.hashKey, bucket.data)
			bucket = bucket.next
		}
		fn(bucket.hashKey, bucket.data)
	}
}

func (h *HashTable) Lens() int {
	return h.lens
}

//TODO 哈希表扩容
func (h *HashTable) Expand()  {

}