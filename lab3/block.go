package main

type block uint64;

func (this *block) setBit(val bool, pos byte) {
	if val {
		*this |= (1 << pos)
	} else {
		*this &= ^(1 << pos)
	}
}

func  (this *block) getBit(pos byte) (bool) {
	val := *this & (1 << pos)
    return (val > 0)
} 

func (this *block) shuffleBits(pos []byte, size byte) {
	shufData := block(0)
	for key, val := range pos {
		shufData.setBit(this.getBit(val), byte(key))
	}
	*this = shufData

	if size != 64 {
		*this &= 1 << size - 1
	}
}

func (this *block) splitCD() (c block, d block) {
	c = *this
	c &= 1 << 28 - 1

	d = (*this) >> 28
	d &= 1 << 28 - 1
	return
}

func MergeCD (c block, d block) block {
	res := c | (d << 28)
	return res
}

func (this *block) shiftL28(shift byte) {
	*this = *this << block(shift)
	*this |= *this >> 28
	*this &= 1 << 28 - 1
}
