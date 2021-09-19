package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

const confPath = "enigma.config"

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
    defer file.Close()
	
	var size int64 = int64(PosN) // stats.Size()
	bytes := make([]byte, size)
    bufr := bufio.NewReader(file)
    // _,err = bufr.Read(bytes)
	
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


func main() {
	rand.Seed(time.Now().UnixNano())
	
	// WriteRndSwapers(fPath, 8)
	swapers, _ := ReadSwapers(confPath)

	rotPick := [...]int{1, 2, 3}
	shiftPick := [...]int{0, 0, 0}
	enigma := NewEnigma(swapers, rotPick, shiftPick)

	var encName string;
	fmt.Println("Путь файла для обработки:")
	fmt.Scanf("%s", &encName)

	inBytes, _ := ioutil.ReadFile(encName)
	outBytes := enigma.encryptArr(inBytes)

	var outName string
	if encName[len(encName)-4:] == ".enc" {
		outName = encName[:len(encName)-4]
	} else {
		outName = fmt.Sprintf("%s.enc", encName)
	}
	fmt.Println("Output file:", outName)

	ioutil.WriteFile(outName, outBytes, 0644)
	return
}
