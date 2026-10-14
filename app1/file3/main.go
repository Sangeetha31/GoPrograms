package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("C:/MyData/Data.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer file.Close()

	_, err1 := file.WriteString("This is Line-1")
	if err1 != nil {
		fmt.Println("Error While Writing Contents:", err1)
	}

	fmt.Println("Write Completed")
}
