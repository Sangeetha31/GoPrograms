package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("userdata.txt")

	if err != nil {
		fmt.Println("Error While Accessing the File:", err)
		return
	}
	file.Close()

	fmt.Println("File Opened Successfully!")

}
