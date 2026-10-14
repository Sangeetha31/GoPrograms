package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:/MyData/Data.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while Reading the file:", err)
	}
}
