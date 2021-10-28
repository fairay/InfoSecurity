package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func IsPrime(n uint64) bool {
	if n % 2 == 0 && n != 2 {
		return false
	}

	for i := uint64(3); i <= uint64(math.Sqrt(float64(n))); i+=2 {
		if n % i == 0 {
			return false;
		}
	}

	return true;
}

func PrimeN(begin uint64, end uint64) uint64 {
	var a uint64 = 4
	for !IsPrime(a) {
		a = (rand.Uint64() % (end-begin)) + begin
	}
	return a
}

func GCD(a uint64, b uint64) uint64 {
	for a * b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	return a + b
}

func ExtGCD(a *big.Int, b *big.Int) (r *big.Int, x *big.Int, y *big.Int) {
	if a.Cmp(big.NewInt(0)) == 0 {
		return b, big.NewInt(0), big.NewInt(1)
	}

	r, x1, y1 := ExtGCD(big.NewInt(0).Mod(b, a), a)
	x = y1.Sub(y1,  big.NewInt(0).Mul(x1, big.NewInt(0).Div(b, a)))
	y = x1
	return
}

func EulersF(p uint64, q uint64) uint64 {
	return (p-1) * (q-1)
}

func GenerateKeys(phi uint64) (pub uint64, pri uint64) {
	pri = 0
	for pri == 0 {
		pub = PublicKey(phi)
		pri = PrivateKey(pub, phi)
	}
	return
}

func PublicKey(phi uint64) (pub uint64) {
	pub = phi
	for GCD(pub, phi) != 1 {
		pub = rand.Uint64() % phi
	}
	return
}

func PrivateKey(pub uint64, phi uint64) uint64 {
	Bpub := big.NewInt(0).SetUint64(pub)
	Bphi := big.NewInt(0).SetUint64(phi)
	_, x, _ := ExtGCD(Bpub, Bphi)

	if x.Cmp(big.NewInt(0)) == -1 {
		return 0
	} else {
		return x.Uint64()
	}
}

func RndKeys() (n, pub, pri uint64) {
	begin := uint64(1 << 31)
	end := uint64(1 << 32) - 1

	p := PrimeN(begin, end)
	q := PrimeN(begin, end)

	n = p*q
	phi := EulersF(p, q)

	pub, pri = GenerateKeys(phi)
	fmt.Println()
	return
}
