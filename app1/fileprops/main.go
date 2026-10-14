package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filepath := "C:/MyData/Data.txt"

	info, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	absPath, err := filepath.Abs(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Name:", info.Name())
	fmt.Println("Absolute Patch:", absPath)
	fmt.Println("Dir Path:", filepath.Dir(absPath))
	fmt.Println("File Extension:", filepath.Ext(info.Name()))
	fmt.Println("File Size:", info.Size(), " Bytes")

	fmt.Println("It is a Folder:", info.IsDir())
	fmt.Println("File Mode:", info.Mode().String())
	fmt.Println("File Permissions:", info.Mode().Perm().String())
	fmt.Println("File Mode(Numeric):", fmt.Sprintf("%04o", info.Mode().Perm()))
	fmt.Println("is Hidden:", len(info.Name()) > 0 && info.Name()[0] == '.')

	fmt.Println("Last Modified:", info.ModTime().Format(time.RFC3339))
	fmt.Println("Last Modified(Local):", info.ModTime().Local().Format(time.RFC3339))

	fmt.Println("Base Name:", filepath.Base(absPath))
	fmt.Println("Parent Dir:", filepath.Dir(filepath.Dir(absPath)))
	fmt.Println("Path Seperator:", string(os.PathListSeparator))

	fmt.Printf("Meta Data: %+v\n", info.Sys())
}
