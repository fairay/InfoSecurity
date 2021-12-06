package main

import (
	"crypto/rsa"
	"crypto/x509"
	"io"
	"os"
)

func KeyPair(rnd io.Reader) (priv *rsa.PrivateKey, pub *rsa.PublicKey) {
	priv, err := rsa.GenerateKey(rnd, 2048)
	if err != nil {
		panic(err)
	}
	pub = &priv.PublicKey
	return
}

func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	arr := x509.MarshalPKCS1PrivateKey(priv)
	return arr
}

func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	arr, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		panic(err)
	}
	return arr
}

func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	key, err := x509.ParsePKCS1PrivateKey(priv)
	if err != nil {
		panic(err)
	}
	return key
}

func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	ifc, err := x509.ParsePKIXPublicKey(pub)
	if err != nil {
		panic(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		panic(err)
	}
	return key
}

func ReadFile(fPath string) []byte {
	f, err := os.Open(fPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 2048)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf[:n]
}

func ReadPublic(fPath string) *rsa.PublicKey {
	buf := ReadFile(fPath)
	return BytesToPublicKey(buf)
}

func ReadPrivate(fPath string) *rsa.PrivateKey {
	buf := ReadFile(fPath)
	return BytesToPrivateKey(buf)
}

func ReadSign(fPath string) []byte {
	return ReadFile(fPath)
}
