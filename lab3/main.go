package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const BlockN = 8;

const keyPath = "1.key"

func EncFile(SrcPath string, DstPath string, KeyPath string) error {
	fSrc, err := os.Open(SrcPath)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fSrc.Close()

	fDst, err := os.Create(DstPath)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fDst.Close()

	key, err := ReadKey(KeyPath)
	roundKeys := GetRoundKeys(key)

	for {
		data, n, err := ReadBlock(fSrc)

		if n != 0 {
			data = EncBlock(data, roundKeys)
			WriteBlock(fDst, data)
		}

		if err == io.EOF {
			s := BlockN - n
			if s == BlockN {
				s = 0
			}
			WriteAddSize(fDst, s)
			break
		} else if err != nil {
			log.Fatal("read failed")
		}
	}

	return nil
}

func DecFile(SrcPath string, DstPath string, KeyPath string) error {
	fSrc, err := os.Open(SrcPath)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fSrc.Close()

	fDst, err := os.Create(DstPath)
	if err != nil {
		log.Fatal("open failed")
	}
	defer fDst.Close()

	key, err := ReadKey(KeyPath)
	roundKeys := GetRoundKeys(key)
	addSize := 0

	for {
		data, n, err := ReadBlock(fSrc)

		if n == BlockN {
			data = DecBlock(data, roundKeys)
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
	rand.Seed(time.Now().UnixNano())

	var sourceFile string;
	switch len(os.Args) {
	case 1:
		log.Fatal("Недостаточно агрументов")
	case 2:
		switch os.Args[1] {
		case "--new-key":
			WriteRndKey(keyPath)
			fmt.Println("Новый ключ создан")
			return
		default:
			sourceFile = os.Args[1]
		}
	default:
		log.Fatal("Неизвестные параметры")
	}

	destFile, isEnc := OutName(sourceFile);
	if isEnc {
		EncFile(sourceFile, destFile, keyPath)
	} else {
		DecFile(sourceFile, destFile, keyPath)
	}
}