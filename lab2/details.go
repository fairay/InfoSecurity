package main

import  (
	"fmt"
	"math/rand"
)
const PosN int = 256;

type SwapI interface {
	forwardTrans(pos byte, offset byte) byte
	backwardTrans(pos byte, offset byte) byte
}

type Swaper struct {
	transArr [PosN]byte
}

func RndSwaper() (s *Swaper) {
	s = new(Swaper)
	for i, el := range rand.Perm(PosN) {
		s.transArr[i] = byte(el)
	}

	return
}

func ByteSwaper(arr []byte) (s *Swaper) {
	s = new(Swaper)
	copy(s.transArr[:], arr[:PosN]);
	return
}

func (this *Swaper) String() (s string) {
	for _, val := range this.transArr {
		s += fmt.Sprintf("%c", val)
		fmt.Printf("%d = \\u%d \n", val, val);
	}
	return
}

func (this *Swaper) forwardTrans(pos byte) byte {
	return this.transArr[int(pos)%PosN]
}

func (this *Swaper) backwardTrans(pos byte) byte {
	for i, val := range this.transArr {
		if int(val) == int(pos)%PosN {
			return byte(i)
		}
	}
	return byte(0)
}


type Rotor struct {
	Swaper 
	shift int
}

func RndRotor() (s *Rotor) {
	s = new(Rotor)
	s.transArr = RndSwaper().transArr
	s.shift = rand.Int() % PosN
	return
}

func (this *Rotor) rotate() bool {
	this.shift = (this.shift + 1) % PosN
	return this.shift == 0
}

func (this *Rotor) forwardTrans(pos byte) byte {
	return this.transArr[(int(pos) + this.shift) % PosN]
}

func (this *Rotor) backwardTrans(pos byte) byte {
	for i, val := range this.transArr {
		if int(val) == (int(pos) + this.shift) % PosN {
			return byte(i)
		}
	}
	return byte(0)
}



type Reflector struct {
	Swaper
}

func RndReflector() (s *Reflector) {
	s = new(Reflector)
	for i, val := range rand.Perm(PosN/2) {
		s.transArr[i] = byte(val) + byte(PosN/2)
	}
	return
}
