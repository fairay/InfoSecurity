package main

// import "fmt"

const RotorN = 3

type SwapConfig [RotorN + 1][PosN]byte

type Enigma struct {
	ref *Reflector
	rot [RotorN]*Rotor
}

func NewEnigma(swapers [][PosN]byte, rotPick [RotorN]int, shiftPick [RotorN]int) (s *Enigma) {
	s = new(Enigma)
	s.ref = new(Reflector)
	s.ref.transArr = swapers[0]

	for i, val := range rotPick {
		s.rot[i] = newRotor(swapers[val], shiftPick[i])
		// s.rot[i].transArr = swapers[val]
		// s.rot[i].shift = shiftPick[i]
	}

	return
}

func (this *Enigma) encryptByte(s byte) byte {
	// fmt.Printf("%d-> ", s);
	for _, val := range this.rot {
		s = val.forwardTrans(s)
		// fmt.Printf("%d -> ", s);
	}

	s = this.ref.forwardTrans(s)
	// fmt.Printf("\t||%d|| -> ", s);

	for i := RotorN - 1; i >= 0; i-- {
		s = this.rot[i].backwardTrans(s)
		// fmt.Printf("%d -> ", s);
	}

	for _, val := range this.rot {
		if !val.rotate() {
			break
		}
	}

	// fmt.Printf("\t%d (%d, %d, %d)\n", s, this.rot[0].shift, this.rot[1].shift, this.rot[2].shift);
	return s
}

func (this *Enigma) encryptArr(sArr []byte) []byte {
	dArr := make([]byte, len(sArr))
	for i, val := range sArr {
		dArr[i] = this.encryptByte(val)
	}
	return dArr
}
