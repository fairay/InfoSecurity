package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type block uint64;
const BlockN = 8;


func EncFile(SrcFName string, DstFName string) error {
	fSrc, err := os.Open(SrcFName)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fSrc.Close()

	fDst, err := os.Create(DstFName)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fDst.Close()

	for {
		data, n, err := ReadBlock(fSrc)

		if n != 0 {
			WriteBlock(fDst, data)
		}

		if err == io.EOF {
			WriteAddSize(fDst, BlockN - n)
			break
		} else if err != nil {
			log.Fatal("read failed")
		}
	}

	return nil
}

func DecFile(SrcFName string, DstFName string) error {
	fSrc, err := os.Open(SrcFName)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fSrc.Close()

	fDst, err := os.Create(DstFName)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fDst.Close()

	addSize := 0
	for {
		data, n, err := ReadBlock(fSrc)

		if n == BlockN {
			err = WriteBlock(fDst, data)
		} else if err == io.EOF && n == 1 {
			addSize = GetAddSize(data)
			break
		} else {
			log.Fatal("read failed")
		}
	}

	fi, err := fDst.Stat()
    if err != nil {
        return err
    }

    return fDst.Truncate(fi.Size() - int64(addSize))
}

func OutName(sName string) (outName string, isEnc bool) {
	if sName[len(sName)-4:] == ".enc" {
		outName = sName[:len(sName)-4]
		isEnc = false
	} else {
		outName = fmt.Sprintf("%s.enc", sName)
		isEnc = true
	}
	return
}

func main() {

	var sourceFile string;
	switch len(os.Args) {
	case 1:
		log.Fatal("Недостаточно агрументов")
	case 3:
		switch os.Args[2] {
		case "-rebuild":
			// regenerate key
		}
	default:
		sourceFile = os.Args[1]
	}

	destFile, isEnc := OutName(sourceFile);
	if isEnc {
		EncFile(sourceFile, destFile)
	} else {
		DecFile(sourceFile, destFile)
	}
}