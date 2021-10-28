package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const priKeyPath = "pri.key"
const pubKeyPath = "pub.key"

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

	n, key, err := ReadKeys(KeyPath)

	for {
		data, nReaded, err := ReadBlock(fSrc, 7)

		if nReaded != 0 {
			data = EncBlock(data, n, key)
			WriteBlock(fDst, data, 8)
		}

		if err == io.EOF {
			s := 7 - nReaded
			if s == 7 {
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

	n, key, err := ReadKeys(KeyPath)
	addSize := 0

	for {
		data, nReaded, err := ReadBlock(fSrc, 8)

		if nReaded == 8 {
			data = DecBlock(data, n, key)
			err = WriteBlock(fDst, data, 7)
		} else if err == io.EOF && nReaded == 1 {
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

	var sourceFile string
	switch len(os.Args) {
	case 1:
		log.Fatal("Недостаточно агрументов")
	case 2:
		switch os.Args[1] {
		case "--new-key":
			WriteRndKeys(pubKeyPath, priKeyPath)
			fmt.Println("Новый ключ создан")
			return
		default:
			sourceFile = os.Args[1]
		}
	default:
		log.Fatal("Неизвестные параметры")
	}

	destFile, isEnc := OutName(sourceFile)
	if isEnc {
		EncFile(sourceFile, destFile, pubKeyPath)
	} else {
		DecFile(sourceFile, destFile, priKeyPath)
	}
}
