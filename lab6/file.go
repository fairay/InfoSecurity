package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFileBytes(fPath string) ([]byte, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	return data, err
}

func WriteFileBytes(fPath string, arr []byte) error {
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(arr)
	return err
}

func ReadFileMap(fPath string) (*CmpMap, error) {
	data, err := ReadFileBytes(fPath)
	if err != nil {
		return nil, err
	}

	m := new(CmpMap)
	err = json.Unmarshal(data, m)
	return m, err
}

func WriteFileMap(fPath string, m *CmpMap) error {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return WriteFileBytes(fPath, data)
}
