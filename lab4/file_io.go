package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadKeys(fName string) (n uint64, key uint64, err error) {
	f, err := os.Open(fName)
	if err != nil {
		log.Fatal("open failed")
	}
	defer f.Close()

	buf := make([]byte, 16)
	nRead, err := f.Read(buf)

	if err != nil || nRead != 16 {
		return 0, 0, errors.New("key corruted")
	}

	n = uint64(binary.BigEndian.Uint64(buf[:8]))
	key = uint64(binary.BigEndian.Uint64(buf[8:]))
	return n, key, nil
}

func WriteBlocks(fName string, blocks []uint64) error {
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal("open failed")
	}
	defer f.Close()

	for _, v := range blocks {
		err := WriteBlock(f, v)
		if err != nil{
			return errors.New("key write failed")
		} 
	}
	return nil
}

func WriteRndKeys(pubF, priF string) (err error) {
	n, pub, pri := RndKeys()
	fmt.Println(n, pub, pri)
	err = WriteBlocks(pubF, []uint64{n, pub})
	if err != nil { return err }
	err = WriteBlocks(priF, []uint64{n, pri})
	return
}

func ReadBlock(f *os.File) (uint64, byte, error) {
	var res uint64
	buf := make([]byte, 8)
	n, err := f.Read(buf)

	if err != nil {
		return 0, 0, err
	}

	res = uint64(binary.BigEndian.Uint64(buf))
	if n != 8 {
		err = io.EOF
	}

	return res, byte(n), err
}

func WriteBlock(f *os.File, data uint64) error {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(data))
	n, err := f.Write(buf)

	if err != nil || n != 8 {
		return errors.New("write fail")
	} else {
		return nil
	}
}

func WriteAddSize(f *os.File, size byte) (err error) {
	buf := make([]byte, 1)
	buf[0] = size
	n, err := f.Write(buf)

	if err != nil || n != 8 {
		return errors.New("write fail")
	} else {
		return nil
	}
}

func GetAddSize(data uint64) (size int) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(data))
	size = int(buf[0])
	return
}
