package main

import (
	"fmt"
	"os"
)

func main() {
	contents := "This is Line-1\nThis is Line-2\nThis is Line-3\nThis is Line-4\n"

	err := os.WriteFile("C:/MyData/Data.txt", []byte(contents), 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Write Operation Completed")
}
