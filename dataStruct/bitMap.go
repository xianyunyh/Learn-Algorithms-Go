package dataStruct

type BitMap struct {
	items []uint
	max   uint
}

//创建一个bitMap
func NewBitMap(max uint) *BitMap {
	lens := max / 32 + 1
	return &BitMap{
		items: make([]uint,lens),
		max:   max,
	}
}
//加入元素
func (b *BitMap)Add(n uint) bool  {
	if b.Exist(n) {
		return false
	}
	//n := b.hashFun(s)
	index := b.getIndex(n)
	b.items[index] =  b.items[index]  | ( 1 << ( n & 0x1F))
	return true
}

func (b *BitMap)Exist(n uint) bool  {
	//n := b.hashFun(s)
	index :=b.getIndex(n)
	pos := b.getYindex(n)
	return (b.items[index] & pos) != 0
}

//获取横坐标a[i]
func (b *BitMap) getIndex(n uint) uint  {
	return n >> 5
}
//获取纵向坐标 a[i][j]
func (b *BitMap) getYindex(n uint) uint {
	return 1 << ( n & 0x1F)
}

//将字符串转成对应的uint
func (b *BitMap) hashFun(s string) uint {
	var sum = rune(0)
	for i := 0;i < len(s);i++ {
		sum += rune(s[i])
	}
	return uint(sum)
}