package main

import (
	"fmt"
	"os"
)

const folderPath = "temp/"
const mapPath = "cmp.map"

func HuffmanEncode(data []byte) ([]byte, *CmpMap) {
	stat := NewStatMap()
	for _, v := range data {
		stat.Add(v)
	}

	tree := NewPrefixTree(stat) // tree.root.Print("")
	tableStruct := tree.ToMap()
	table := tableStruct.Table

	cmpBits := EmptyBitSet()
	for _, v := range data {
		cmpBits.Add(table[v])
	}

	copyBytes := cmpBits.ToBytes()
	return copyBytes, tableStruct
}

func HuffmanDecode(data []byte, m *CmpMap) []byte {
	cmpBits := BitSetFromBytes(data)
	tempBits := EmptyBitSet()

	ans := make([]byte, 0)
	im := m.ITable()
	for i := cmpBits.Len - 1; i >= 0; i-- {
		tempBits.ForwardBit(cmpBits.Val(i))

		if val, found := im.FindVal(tempBits); found {
			ans = append(ans, val)
			tempBits = EmptyBitSet()
		}
	}
	
	// Reverse ans
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return []byte(ans)
}

func main() {
	switch len(os.Args) {
	case 1:
		panic(fmt.Errorf("Недостаточно агрументов"))
	case 2:
		panic(fmt.Errorf("Недостаточно агрументов"))
	case 3:
		panic(fmt.Errorf("Недостаточно агрументов"))
	case 4:
		switch os.Args[1] {
		case "--cmp":
			// Example: .\lab6.exe --cmp parrot.jpg parrot.cmp
			data, err := ReadFileBytes(folderPath + os.Args[2])
			if err != nil {
				panic(err)
			}

			data, m := HuffmanEncode(data)

			WriteFileBytes(folderPath+os.Args[3], data)
			WriteFileMap(folderPath+mapPath, m)
		case "--uncmp":
			// Example: .\lab6.exe --uncmp parrot.cmp new.jpg
			data, err := ReadFileBytes(folderPath + os.Args[2])
			if err != nil {
				panic(err)
			}

			m, err := ReadFileMap(folderPath+mapPath)
			if err != nil {
				panic(err)
			}

			data = HuffmanDecode(data, m)
			WriteFileBytes(folderPath+os.Args[3], data)
		default:
			panic(fmt.Errorf("Неизвестная комманда"))
		}

	default:
		panic(fmt.Errorf("Неизвестные параметры"))
	}
}
