package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("Go/Demo/Sample", 0755)

	if err != nil {
		fmt.Println("Error While creating folders", err)
		return
	}

	fmt.Println("Folder(s) Created Successfully")
}
