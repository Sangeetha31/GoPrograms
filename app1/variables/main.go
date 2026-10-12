package main

import "fmt"

func main() {
	var a int = 10

	var b int = 20

	b = b + 1
	fmt.Println("The value of a is:", a)
	fmt.Println("The value of b is:", b)

	fmt.Printf("The value of a is: %d", a)
	fmt.Printf("The value of b is: %d", b)
}
