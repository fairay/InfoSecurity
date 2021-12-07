package main

type BitSet struct {
	Len int      `json:"len"`
	Arr []uint64 `json:"arr"`
}

type CmpMap struct {
	Table map[byte]*BitSet `json:"table"`
}

func EmptyBitSet() *BitSet {
	obj := &BitSet{
		Len: 0,
		Arr: make([]uint64, 1),
	}
	return obj
}
func NewBitSet(len int, num uint64) *BitSet {
	obj := &BitSet{
		Len: len,
		Arr: make([]uint64, 1),
	}
	obj.Arr[0] = num
	return obj
}
func BitSetFromBytes(arr []byte) *BitSet {
	this := EmptyBitSet()
	addSize := arr[len(arr)-1]
	arr = arr[:len(arr)-1]

	for _, v := range arr {
		this.Addb(8, uint64(v))
	}
	this.Len = len(arr)*8 - int(addSize)
	return this
}
func (this *BitSet) Copy() *BitSet {
	obj := &BitSet{
		Len: this.Len,
		Arr: make([]uint64, (this.Len+64)/64),
	}
	copy(obj.Arr, this.Arr)

	return obj
}

func (this *BitSet) ToBytes() []byte {
	arr := make([]byte, 0)
	for i := 0; i < (this.Len+7)/8; i++ {
		arr = append(arr, this.Byte(i))
	}

	arr = append(arr, byte(8-this.Len%8))
	return arr
}

func (this *BitSet) Addb(len int, data uint64) {
	this.Add(NewBitSet(len, data))
}

func (this *BitSet) Add(other *BitSet) {
	pos := this.Len / 64
	rem := this.Len % 64
	if rem+other.Len <= 64 {
		num := other.Arr[0]
		num = num << rem

		this.Arr[pos] |= num
		this.Len += other.Len
	} else {
		overflow := rem + other.Len - 64

		numB := other.Arr[0] >> (other.Len - overflow)
		numL := other.Arr[0] << rem

		this.Arr[pos] |= numL
		this.Arr = append(this.Arr, numB)
		this.Len += other.Len
	}

	if rem+other.Len == 64 {
		this.Arr = append(this.Arr, 0)
	}
}

func (this *BitSet) ForwardBit(bit bool) {
	add := uint64(0)
	if bit {
		add = 1
	}
	this.Arr[0] = this.Arr[0]<<1 | add
	this.Len += 1
}

func (this *BitSet) Val(pos int) bool {
	i := pos / 64
	val := this.Arr[i] >> (pos % 64)
	return val&1 == 1
}
func (this *BitSet) Byte(pos int) byte {
	i := pos / 8
	off := (pos % 8) * 8
	val := this.Arr[i] >> off
	return byte(val)
}

func (this *CmpMap) FindVal(bits *BitSet) (val byte, found bool) {
	for k, v := range this.Table {
		if bits.Len == v.Len && bits.Arr[0] == v.Arr[0] {
			return k, true
		}
	}
	return 0, false
}

type ICmpMap struct {
	Table map[uint64]map[int]byte
}

func (this *CmpMap) ITable() (m *ICmpMap) {
	m = &ICmpMap{
		Table: make(map[uint64]map[int]byte),
	}

	for k, v := range this.Table {
		if _, ok := m.Table[v.Arr[0]]; ok {
			m.Table[v.Arr[0]][v.Len] = k
		} else {
			m.Table[v.Arr[0]] = map[int]byte{v.Len: k}
		}
	}
	return m
}

func (this *ICmpMap) FindVal(bits *BitSet) (val byte, found bool) {
	if _, ok := this.Table[bits.Arr[0]]; ok {
		val, found = this.Table[bits.Arr[0]][bits.Len]
	} else {
		val, found = 0, false
	}
	return
}
