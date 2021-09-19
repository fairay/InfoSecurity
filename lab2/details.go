package main

import  (
	//"fmt"
	"math/rand"
)
const PosN int = 256;


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
	backTransArr [PosN]byte
	shift int
}

func RndRotor() (s *Rotor) {
	s = new(Rotor)
	s.transArr = RndSwaper().transArr
	s.shift = rand.Int() % PosN
	return
}

func newRotor(arr [PosN]byte, _shift int) *Rotor {
	this := new(Rotor)
	this.transArr = arr
	this.shift = _shift

	for i, val := range arr {
		this.backTransArr[val] = byte(i)
	}

	return this
}

func (this *Rotor) rotate() bool {
	this.shift = (this.shift + 1) % PosN
	return this.shift == 0
}

func (this *Rotor) forwardTrans(pos byte) byte {
	val := this.transArr[(int(pos) + this.shift) % PosN]
	return byte((int(val) - this.shift) % PosN)
}

func (this *Rotor) backwardTrans(pos byte) byte {
	preVal := this.backTransArr[(int(pos) + this.shift) % PosN]
	return byte((int(preVal) - this.shift) % PosN)
}


type Reflector struct {
	Swaper
}

func RndReflector() (s *Reflector) {
	s = new(Reflector)
	for i, val := range rand.Perm(PosN/2) {
		s.transArr[i] = byte(val) + byte(PosN/2)
		s.transArr[s.transArr[i]] = byte(i)
	}
	return
}
