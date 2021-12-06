package main

import (
	"math/big"
)

func powMod(data, pow, n uint64) uint64 {
	d := big.NewInt(1)
	t := big.NewInt(0).SetUint64(data)
	nBig := big.NewInt(0).SetUint64(n)
	for pow != 0 {
		if pow % 2 == 1 {
			d.Mod(d.Mul(d, t), nBig)
		}
		t.Mod(t.Mul(t, t), nBig)
		pow /= 2
	}
	return d.Uint64()
}

func EncBlock(data uint64, n, key uint64) uint64 {
	return powMod(data, key, n)
}

func DecBlock(data uint64, n, key uint64) uint64 {
	return powMod(data, key, n)
}
