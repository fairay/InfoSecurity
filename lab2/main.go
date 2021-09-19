package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	// "io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	// "os/exec"
	"time"
	// "ioutil"
	// "encoding/binary"
)

const fPath = "enigma.config"

func WriteRndSwapers(filename string, rotN int) (error) {
	f, err := os.Create(filename)
    if err != nil {
        log.Fatal("Couldn't open file")
    }
    defer f.Close()

	data := RndReflector().transArr
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


func ReadFile(filename string)  ([]byte, error) {
	return nil, nil
}

func WriteFile(filename string, data []byte) (error) {
	return nil
}


func RetrieveROM(filename string) ([]byte, error) {
    file, err := os.Open(filename)

    if err != nil {
        return nil, err
    }
    defer file.Close()

    stats, statsErr := file.Stat()
    if statsErr != nil {
        return nil, statsErr
    }
	fmt.Println("file size ", stats.Size()) //

    var size int64 = int64(PosN) // stats.Size()
    bytes := make([]byte, size)

    bufr := bufio.NewReader(file)
    // _,err = bufr.Read(bytes)
	_, err = io.ReadFull(bufr, bytes)
	_, err = io.ReadFull(bufr, bytes)

    return bytes, err
}

func ReadStruct(filename string) (enigma *Enigma, err error) {
	file, err := os.Open(filename)
    defer file.Close()

	var swapers [][PosN]byte;
    binary.Read(file, binary.BigEndian, swapers)
// var hi struct {
	// 	S1 [PosN]byte
	// 	S2 [PosN]byte
	// }
	return
}

func WriteStruct(filename string, enigma *Enigma) (error) {
	f, err := os.Create("file.bin")
    if err != nil {
        log.Fatal("Couldn't open file")
    }
    defer f.Close()

	data := RndSwaper().transArr
	for _, val := range data {
		fmt.Printf("%d ", val);
	}

    err = binary.Write(f, binary.BigEndian, data)
    if err != nil {
        log.Fatal("Write failed")
    }	

	return err;
}


func main() {
	rand.Seed(time.Now().UnixNano())
	
	WriteRndSwapers(fPath, 8)
	swapers, _ := ReadSwapers(fPath)
	// fmt.Println("\njopa\n", swapers)

	rotPick := [...]int{3, 4, 1}
	shiftPick := [...]int{91, 21, 7}
	enigma := ByteEnigma(swapers, rotPick, shiftPick)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Путь файла для шифрования:")
	encName, _ := reader.ReadString('\n')
	fmt.Println(encName)

	inBytes, _ := ioutil.ReadFile(encName)
	outBytes := enigma.encryptArr(inBytes)
	ioutil.WriteFile(encName+".enc", outBytes, 0644)
	return

	/*
	f, err := os.Create("file.bin")
    if err != nil {
        log.Fatal("Couldn't open file")
    }
    defer f.Close()

	data := RandomSwaper().transArr
	for _, val := range data {
		fmt.Printf("%d ", val);
	}

    err = binary.Write(f, binary.BigEndian, data)
    if err != nil {
        log.Fatal("Write failed")
    }
	*/

	// swap := RandomSwaper()
	// fmt.Print(swap);

	en := RndEnigma()
	WriteStruct("file.bin", en)
	fmt.Println(en)

	fmt.Println("\n\n+++\n\n")

	bts, _ := ReadStruct("file.bin") //RetrieveROM("file.bin");
	print(bts)
	for _, val := range bts.ref.transArr {
		fmt.Printf("%d ", val)
	}
	
	// fmt.Printf("%c\n", swap.forwardTrans('a'))
}
