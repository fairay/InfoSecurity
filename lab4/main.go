package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func IsPrime(n uint64) bool {
	if n % 2 == 0 && n != 2 {
		return false
	}

	for i:=uint64(3); i <= uint64(math.Sqrt(float64(n))); i+=2 {
		if n % i == 0 {
			return false;
		}
	}

	return true;
}

func PrimeN(begin uint64, end uint64) uint64 {
	var a uint64
	for {
		a = (rand.Uint64() % (end-begin)) + begin
		if IsPrime(a) {
			break
		}
	}
	return a
}

func GCD(a uint64, b uint64) uint64 {
	for a * b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	return a+b
}

func ExtGCD(a uint64, b uint64) (r uint64, x uint64, y uint64) {
	if a == 0 {
		return b, 0, 1;
	}

	r, x1, y1 := ExtGCD(b%a, a);
	x = y1 - (b / a) * x1;
	y = x1;
	return
}

func EulersF(p uint64, q uint64) uint64 {
	return (p-1) * (q-1)
}

func PublicKey(phi uint64) uint64 {
	begin := uint64(1 << 31)
	end := uint64(1 << 32) - 1
	var a uint64
	for {
		a = (rand.Uint64() % (end-begin)) + begin
		if GCD(a, phi) == 1 {
			break
		}
	}
	return a
}

func PrivateKey(pub uint64, phi uint64) uint64 {
	return 0
}

func main() {
	rand.Seed(time.Now().UnixNano())

	begin := uint64(1 << 31)
	end := uint64(1 << 32) - 1

	p := PrimeN(begin, end)
	q := PrimeN(begin, end)

	N := p*q
	phi := EulersF(p, q)


	_, x, _ := ExtGCD(11, 71)
	fmt.Println(math.MaxUint64(x, 0-x))
	fmt.Println(p, q)
	fmt.Println(N)
	fmt.Println(phi)
	fmt.Println(PublicKey(phi))
}
