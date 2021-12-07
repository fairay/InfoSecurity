package main

type bitSet struct {
	len int
	arr []uint64
}

type cmpMap struct {
	table map[byte]*bitSet
}

func EmptyBitSet() *bitSet {
	obj := &bitSet{
		len: 0,
		arr: make([]uint64, 1),
	}
	return obj
}
func NewBitSet(len int, num uint64) *bitSet {
	obj := &bitSet{
		len: len,
		arr: make([]uint64, 1),
	}
	obj.arr[0] = num
	return obj
}
func BitSetFromBytes(arr []byte) *bitSet {
	this := EmptyBitSet()
	addSize := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	
	for _, v := range arr {
		this.Addb(8, uint64(v))
	}
	this.len = len(arr)*8 - int(addSize)
	return this
}
func (this *bitSet) Copy() *bitSet {
	obj := &bitSet{
		len: this.len,
		arr: make([]uint64, (this.len+64)/64),
	}
	copy(obj.arr, this.arr)

	return obj
}


func (this *bitSet) ToBytes() []byte{
	arr := make([]byte, 0)
	for i := 0; i < (this.len + 7) / 8; i++ {
		arr = append(arr, this.Byte(i))
	}

	arr = append(arr, byte(8 - this.len % 8))
	return arr
}

func (this *bitSet) Addb(len int, data uint64) {
	this.Add(NewBitSet(len, data))
}

func (this *bitSet) Add(other *bitSet) {
	pos := this.len / 64
	rem := this.len % 64
	if rem+other.len <= 64 {
		num := other.arr[0]
		num = num << rem

		this.arr[pos] |= num
		this.len += other.len
	} else {
		overflow := rem + other.len - 64

		numB := other.arr[0] >> (other.len - overflow)
		numL := other.arr[0] << rem

		this.arr[pos] |= numL
		this.arr = append(this.arr, numB)
		this.len += other.len
	}

	if rem+other.len == 64 {
		this.arr = append(this.arr, 0)
	}
}

func (this *bitSet) BackBit(bit bool) {
	if bit {
		this.arr[0] = this.arr[0] | (1 << (this.len))
	}
	this.len += 1
	// this.Add(NewBitSet(1, add))
}

func (this *bitSet) ForvardBit(bit bool) {
	add := uint64(0)
	if bit {
		add = 1
	}
	this.arr[0] = this.arr[0]<<1 | add
	this.len += 1
}

func (this *bitSet) Val(pos int) bool {
	i := pos / 64
	val := this.arr[i] >> (pos % 64)
	return val&1 == 1
}
func (this *bitSet) Byte(pos int) byte {
	i := pos / 8
	off := (pos % 8)*8
	val := this.arr[i] >> off
	return byte(val)
}


func (this *cmpMap) FindVal(bits *bitSet) (val byte, found bool) {
	for k, v := range this.table {
		if bits.len == v.len && bits.arr[0] == v.arr[0] {
			return k, true
		}
	}
	return 0, false
}
