package main

import (
	"des/tab"
	"fmt"
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
		fmt.Println(k[i])
		k[i].shuffleBits(tab.CP, 48)
	}
	return
}
