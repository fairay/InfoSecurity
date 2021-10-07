package main

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

func ReadBlock(f *os.File) (block, byte, error) {
	var res block
	buf := make([]byte, BlockN)
	n, err := f.Read(buf)
	
	if err != nil {
		return 0, 0, err
	}

	res = block(binary.BigEndian.Uint64(buf))
	if n != BlockN {
		err = io.EOF
	}

	return res, byte(n), err
}

func WriteBlock(f *os.File, data block) (error) {
	buf := make([]byte, BlockN)
	binary.BigEndian.PutUint64(buf, uint64(data))
	n, err := f.Write(buf)

	if err != nil || n != BlockN {
		return errors.New("write fail")
	} else {
		return nil
	}
}

func WriteAddSize(f *os.File, size byte) (err error) {
	buf := make([]byte, 1)
	buf[0] = size
	n, err := f.Write(buf)

	if err != nil || n != BlockN {
		return errors.New("write fail")
	} else {
		return nil
	}
}

func GetAddSize(data block) (size int) {
	buf := make([]byte, BlockN)
	binary.BigEndian.PutUint64(buf, uint64(data))
	size = int(buf[0])
	return
}



