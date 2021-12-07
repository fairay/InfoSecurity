package main

import (
	"io/ioutil"
	"os"
)

func ReadFileBytes(fPath string) ([]byte, error) {
	f, err := os.Open("test.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFileBytes(fPath string, arr []byte) error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(arr)
	return err
}
