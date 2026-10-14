package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	dirPath := "./Go"

	entries, err := os.ReadDir(dirPath)

	if err != nil {
		fmt.Println("Error While Reading Folder:", err)
		return
	}

	for _, entry := range entries {
		fullpath := filepath.Join(dirPath, entry.Name())

		info, err := entry.Info()

		if err != nil {
			fmt.Println("Error To Read File Props:", err)
			continue
		}
		fmt.Println("Name:", info.Name())
		fmt.Println("Full Path:", fullpath)
		fmt.Println("Size:", info.Size(), " Bytes")
		fmt.Println("-----------------------")
	}
}
