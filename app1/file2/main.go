package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "C:/MyData/Data.txt"

	err := os.WriteFile(filename, []byte("Sample Text Here"), 0644)

	if err != nil {
		fmt.Println("Error Occured:", err)
		return
	}

	fmt.Println("Write Operation Completed Successfully!")
	//4 - Read
	//2- Write
	//1- Execute
	//0- No Permission
}
