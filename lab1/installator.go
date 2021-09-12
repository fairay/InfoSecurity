package main

import (
	"fmt"
)

func mainInstallator() {
	err := fwriteKey("key");
	if err != 0 {
		fmt.Println("Key writing error");
	} else {
		fmt.Println("Key successfully writen, enjoy your program");
	}
}

