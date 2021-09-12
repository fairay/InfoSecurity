package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

func getUuid() string {
	bite_res, _ := exec.Command("cmd", "/C", "wmic", "csproduct", "get", "uuid").Output();
	var res string;
	res = strings.Split(string(bite_res), "\n")[1];
	return res;
}

func getHashedKey() string {
	var resStr string;
	byteKey := []byte(getUuid())
	hashedKey := sha256.Sum256(byteKey);
	resStr = fmt.Sprintf("%x", hashedKey);
	return resStr;
}

func freadKey(fName string) (int, string) {
	var res string
	file, err := os.Open(fName)
	if err != nil {
		return -1, ""
	}

	byteRes, err := ioutil.ReadAll(file)
	file.Close()
	
	res = string(byteRes)
    return 0, res;
}

func fwriteKey(fName string) int {
	file, err := os.Create(fName)
	if err != nil {
		return -1;
	}

    mw := io.MultiWriter(os.Stdout, file)
    _, err = fmt.Fprint(mw, getHashedKey())
	file.Close();
	if err != nil {
		return -2;
	}
	return 0;
}
