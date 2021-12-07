package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
)

const priKeyPath = "pri.key"
const pubKeyPath = "pub.key"

func FileHash(fPath string) []byte {
	f, err := os.Open(fPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	hash := crypto.SHA512.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		panic(err)
	}

	hash.Write([]byte(fPath))
	return hash.Sum(nil)
}

func WriteFile(fName string, blocks []byte) error {
	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(blocks)
	return err
}

func WriteRndKeys(pubF, priF string) error {
	pri, pub := KeyPair(rand.Reader)

	err := WriteFile(pubF, PublicKeyToBytes(pub))
	if err != nil {
		return err
	}

	err = WriteFile(priF, PrivateKeyToBytes(pri))
	return err
}

func CreateSign(sFile string) {
	pub := ReadPublic(pubKeyPath)

	hashedFile := FileHash(sFile)
	chipher, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, pub, hashedFile, []byte(""))
	if err != nil {
		fmt.Println("Ошибка шифрования!")
		return
	}
	WriteFile(fmt.Sprintf("%s.sig", sFile), chipher)

	fmt.Println("Подпись создана")
}

func CheckSign(sFile, signFile string) {
	pri := ReadPrivate(priKeyPath)
	sign := ReadSign(signFile)
	unchipher, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, pri, sign, []byte(""))
	if err != nil {
		fmt.Println("Подпись некорректна!")
		return
	}
	origHash := FileHash(sFile)

	if string(origHash) == string(unchipher) {
		fmt.Println("Подпись подтверждена")
	} else {
		fmt.Println("Подпись не подтверждена!")
	}
}

func main() {
	switch len(os.Args) {
	case 1:
		log.Fatal("Недостаточно агрументов")
	case 2:
		switch os.Args[1] {
		case "--new-key":
			WriteRndKeys(pubKeyPath, priKeyPath)
			fmt.Println("Новый ключ создан")
			return
		}
		CreateSign(os.Args[1])
		return
	case 3:
		CheckSign(os.Args[1], os.Args[2])

	default:
		log.Fatal("Неизвестные параметры")
	}
}
