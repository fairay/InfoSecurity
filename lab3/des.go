package main

import (
	"des/tab"
	"math/rand"
)

const RoundN = 16


func RndKey() (block) {
	return block(rand.Uint64());
}

func GetRoundKeys(k0 block) (k [RoundN]block) {
	k0.shuffleBits(tab.B, 56)
	c, d := k0.splitCD()
	for i, shift := range tab.Si {
		c.shiftL28(shift)
		d.shiftL28(shift)

		k[i] = MergeCD(c, d)
		k[i].shuffleBits(tab.CP, 48)
	}
	return
}

func FeistelF(mr block, key block) (res block) {
	mr.shuffleBits(tab.E, 48)
	z := mr ^ key
	for i := 7; i >= 0; i-- {
		off := byte(i)*4
		x := z.getBits([]byte{ off+0, off+5 })
		y := z.getBits([]byte{ off+1, off+2, off+3, off+4 })
		
		si := tab.S[i][x][y]
		res.appendB(block(si), 4)
	}
	res.shuffleBits(tab.P, 32)

	return
}

func EncBlock(data block, keys [RoundN]block) (block) {
	data.shuffleBits(tab.IP, 64)
	l, r := data.splitLR()

	for i:=0; i<RoundN; i++ {
		l, r = r, l ^ FeistelF(r, keys[i])
	}
	
	data = MergeLR(l, r)
	data.shuffleBits(tab.NegIP, 64)
	return data
}

func DecBlock(data block, keys [RoundN]block) (block) {
	data.shuffleBits(tab.IP, 64)
	l, r := data.splitLR()

	for i:=RoundN-1; i>=0; i-- {
		r, l = l, r ^ FeistelF(l, keys[i])
	}
	
	data = MergeLR(l, r)
	data.shuffleBits(tab.NegIP, 64)
	return data
}

