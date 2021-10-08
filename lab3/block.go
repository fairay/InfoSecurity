package main

type block uint64

func (this *block) setBit(val bool, pos byte) {
	if val {
		*this |= (1 << pos)
	} else {
		*this &= ^(1 << pos)
	}
}

func (this *block) getBit(pos byte) bool {
	val := *this & (1 << pos)
	return (val > 0)
}

func (this *block) getBits(pos []byte) (res block) {
	for i, val := range pos {
		res.setBit(this.getBit(val), byte(i))
	}
	return res
}

func (this *block) appendB(data block, size byte) {
	*this = *this << block(size)
	*this |= data
}

/*
func (this *block) shuffleBits(pos []byte, size byte) {
	shufData := block(0)
	for key, val := range pos {
		shufData.setBit(this.getBit(val), byte(key))
	}
	*this = shufData

	if size != 64 {
		*this &= 1<<size - 1
	}
}*/

func (this *block) shuffleBits(pos []byte, size byte) {
	shufData := block(0)
	for key, val := range pos {
		shufData.setBit(this.getBit(val-1), byte(key))
	}
	*this = shufData

	if size != 64 {
		*this &= 1<<size - 1
	}
}

func (this *block) splitCD() (c block, d block) {
	c = *this
	c &= 1<<28 - 1

	d = (*this) >> 28
	d &= 1<<28 - 1
	return
}

func (this *block) splitLR() (l block, r block) {
	r = *this
	r &= 1<<32 - 1

	l = (*this) >> 32
	l &= 1<<32 - 1
	return
}

func MergeCD(c block, d block) block {
	res := c | (d << 28)
	return res
}

func MergeLR(l block, r block) block {
	res := r | (l << 32)
	return res
}

func (this *block) shiftL28(shift byte) {
	*this = *this << block(shift)
	*this |= (*this >> 28) & (1<<block(shift) - 1)
	*this &= 1<<28 - 1
}
