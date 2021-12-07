package main

import (
	"fmt"
)

func HuffmanEncode(data []byte) {
	stat := NewStatMap()
	for _, v := range data {
		stat.Add(v)
	}

	tree := NewPrefixTree(stat)
	tree.root.Print("")
	tableStruct := tree.ToMap()
	table := tableStruct.table

	cmpBits := EmptyBitSet()
	for _, v := range data {
		cmpBits.Add(table[v])
	}

	copyBytes := cmpBits.ToBytes()
	WriteFileBytes("ruko.blud", copyBytes)
	fmt.Println(copyBytes)

	copyBytes, _ = ReadFileBytes("ruko.blud")
	fmt.Println(copyBytes)
	cmpBits = BitSetFromBytes(copyBytes)
	
	tempBits := EmptyBitSet()
	ans := ""
	for i := cmpBits.len-1; i >= 0; i-- {
		tempBits.ForvardBit(cmpBits.Val(i))
		if val, found := tableStruct.FindVal(tempBits); found {
			ans = string(val) + ans
			tempBits = EmptyBitSet()
		}
	}
	fmt.Println()
	fmt.Println(ans)
}

func main() {
	// HuffmanEncode([]byte("abcdefghijklmnop"))
	str := "abcdefghijabcdefgagrnclfkabcdefgagrnclfk                    cum"
	HuffmanEncode([]byte(str)) // abcdefgagrnclfk
	fmt.Println(str)
}
