package dataStruct

type PHPHashTable struct {
	TableMask uint64
	NumUsed uint32
	NumElements uint32
	TableSize uint32
	ArData *Buckets
}

type Buckets struct {
	Index []int
	Items []*Bucket
}

func NewPHPHashTable() *PHPHashTable  {
	idx := make([]int,8)
	items := make([]*Bucket,8)
	for i := 0;i < 8;i++  {
		idx = append(idx,-1)
	}
	return  &PHPHashTable{
		TableMask:   0,
		NumUsed:     0,//使用的个数 包含被删除的
		NumElements: 0,//真实的个数
		TableSize:   0,//总数
		ArData:     &Buckets{
			Index: idx,
			Items: items,
		},
	}
}

func (self *PHPHashTable)Insert(key string,value interface{})  {
	idx := Time33(key) % self.TableMask
	//冲突
	bucket := &Bucket{
		data:    value.(string),
		hashKey: key,
	}
	if self.ArData.Index[idx] != -1 {
		arrIdx := self.ArData.Index[idx]
		bucket.next = self.ArData.Items[arrIdx]
		self.ArData.Index[idx] = int(self.NumUsed)
		self.ArData.Items[idx] = bucket
	} else {
		self.ArData.Index[idx] = int(self.NumUsed)
		self.ArData.Items[self.NumUsed] = bucket
	}
	self.NumUsed++
	self.NumElements++
	self.TableSize++
}
func (self *PHPHashTable)Find(key string) interface{}  {
	idx := Time33(key) % self.TableMask
	itemIdx := self.ArData.Index[idx]
	//没找到
	if itemIdx == -1 {
		return nil
	}
	//遍历链表
	item := self.ArData.Items[itemIdx]
	for item.next != nil {
		if item.hashKey == key {
			return item.data
		}
		item = item.next
	}
	return  nil
}

//time 33 算法
func Time33(str string) uint64  {
	var  hashMask uint64  = 5381
	for _,s := range str {
		c := uint64(s)
		hashMask += ( c >> 5) + c // *33 = h×32+ h
	}
	return hashMask & 0x7FFFFFFF
}