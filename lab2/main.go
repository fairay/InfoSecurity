package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
// 	"math"
	"math/rand"
	"os"
	"time"
)

const confFile = "enigma.config"
const keyFile = "1.key";

func WriteRndSwapers(filename string, rotN int) (error) {
	f, err := os.Create(filename)
    if err != nil {
        log.Fatal("Couldn't open file")
    }
    defer f.Close()

	data := RndReflector().transArr
	for _, val := range data {
		fmt.Printf("%d ", val)
	}
	fmt.Println("\n===")
	err = binary.Write(f, binary.BigEndian, data)
	if err != nil {
		log.Fatal("Write failed")
	}	

	for i := 0; i < rotN; i++ {
		data := RndSwaper().transArr
		for _, val := range data {
			fmt.Printf("%d ", val)
		}
		fmt.Println("\n")
		err = binary.Write(f, binary.BigEndian, data)
		if err != nil {
			log.Fatal("Write failed")
		}	
	}

	return err;
}

func ReadSwapers(filename string) (swap_arr [][PosN]byte, err error) {
	file, err := os.Open(filename)
	if err != nil {
        log.Fatal("Couldn't open file")
    }
    defer file.Close()
	
	var size int64 = int64(PosN)
	bytes := make([]byte, size)
    bufr := bufio.NewReader(file)
	
	for {
		n, err := io.ReadFull(bufr, bytes);
		if (err != nil || n != PosN) {
			break
		}
		swap_arr = append(swap_arr, [PosN]byte{})
		copy(swap_arr[len(swap_arr)-1][:], bytes)
	}
	
	if err == io.EOF {
		err = nil	
	} 
	return
}

func ReadKey(filename string) (rotPick [RotorN]int, shiftPick [RotorN]int, err error) {
	file, err := os.Open(filename)
	if err != nil {
        log.Fatal("Couldn't open file")
    }
	defer file.Close()

	for i := 0; i < RotorN; i++ {
		fmt.Fscanf(file, "%d", &rotPick[i])
	}

	for i := 0; i < RotorN; i++ {
		fmt.Fscanf(file, "%d", &shiftPick[i])
	}

	return
}

func BuildEnigma() (e *Enigma, err error) {
	swapers, err := ReadSwapers(confFile)
	if err != nil { return }

	rot, shift, err := ReadKey(keyFile)
	if err != nil { return }

	// rotPick := [...]int{1, 2, 3}
	// shiftPick := [...]int{0, 0, 0}
	e = NewEnigma(swapers, rot, shift)
	return
}

func OutName(sName string) (outName string) {
	if sName[len(sName)-4:] == ".enc" {
		outName = sName[:len(sName)-4]
	} else {
		outName = fmt.Sprintf("%s.enc", sName)
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	var sourceFile string;
	switch len(os.Args) {
	case 1:
		log.Fatal("Недостаточно агрументов")
	case 3:
		switch os.Args[2] {
		case "-rebuild":
			WriteRndSwapers(confFile, 8)
		}
	default:
		sourceFile = os.Args[1]
		fmt.Println("Путь файла для обработки: ")
	}

	enigma, _ := BuildEnigma()

	// fmt.Print("Путь файла для обработки: ")
	// fmt.Scanf("%s", &sourceFile)

	inBytes, _ := ioutil.ReadFile(sourceFile)
	outBytes := enigma.encryptArr(inBytes)

	destFile := OutName(sourceFile);
	fmt.Println("Результирующий файл: ", destFile)
	ioutil.WriteFile(destFile, outBytes, 0644)

	return
}
