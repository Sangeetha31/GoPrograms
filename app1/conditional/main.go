package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Enter 3 Numbers:")
	fmt.Scan(&a, &b, &c)

	if a > b && a > c {
		fmt.Println(a)
	} else if b > c {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}

}
