package main

import "fmt"

const RotorN = 3

type SwapConfig [RotorN + 1][PosN]byte

type Enigma struct {
	ref Reflector
	rot [RotorN]Rotor
}

func RndEnigma() (s *Enigma) {
	s = new(Enigma)
	s.ref = *RndReflector()
	for i := range &s.rot {
		s.rot[i] = *RndRotor()
	}
	return s
}

func ByteEnigma(swapers [][PosN]byte, rotPick [RotorN]int, shiftPick [RotorN]int) (s *Enigma) {
	s = new(Enigma)
	s.ref.transArr = swapers[0]

	for i, val := range rotPick {
		s.rot[i].transArr = swapers[val]
	}

	return
}

func (this *Enigma) encryptByte(s byte) byte {
	fmt.Printf("%d -> ", s);
	for _, val := range this.rot {
		s = val.forwardTrans(s)
	}

	s = this.ref.forwardTrans(s)

	for i := RotorN - 1; i >= 0; i-- {
		s = this.rot[i].backwardTrans(s)
	}

	for _, val := range this.rot {
		if !val.rotate() {
			break
		}
	}

	fmt.Printf("%d (%d, %d, %d)\n", s, this.rot[0].shift, this.rot[1].shift, this.rot[2].shift);
	return s
}

func (this *Enigma) encryptArr(sArr []byte) []byte {
	dArr := make([]byte, len(sArr))
	for i, val := range sArr {
		dArr[i] = this.encryptByte(val)
	}
	return dArr
}
