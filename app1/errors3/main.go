package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Before Panic")
	panic("OOPS Something Went Wrong!")
	fmt.Println("After Panic")

}
