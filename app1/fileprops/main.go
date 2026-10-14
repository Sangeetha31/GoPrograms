package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "C:/MyData/Data.txt"

	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	absPath, err := filePath.Abs(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Name:", info.Name())
	fmt.Println("Absolute Patch:", absPath)
	fmt.Println("Dir Path:", filePath.Dir(absPath))
	fmt.Println("File Extension:", filePath.Ext(info.Name()))
	fmt.Println("File Size:", info.Size(), " Bytes")

	fmt.Println("It is a Folder:", info.IsDir())
	fmt.Println("File Mode:", info.Mode().String())

}
