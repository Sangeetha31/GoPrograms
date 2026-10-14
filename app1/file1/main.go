package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("C:/MyData/Data.txt")

	if err != nil {
		fmt.Println("Error Occured:", err)
	}

	defer file.Close()

	fmt.Println("File Created Successfully!")
}
