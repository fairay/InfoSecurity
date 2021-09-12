package main

import (
	"fmt"
)

func mainApp() {
	var validKey, storedKey string;
	validKey = getHashedKey();
	err, storedKey := freadKey("key");
	if err != 0 {
		fmt.Println("key file is missing")
		return
	} else if validKey != storedKey {
		fmt.Println("key isn't valid")
		return
	}

	fmt.Println("Welcome to the programm");
}

func main() {
	mainApp();
}
