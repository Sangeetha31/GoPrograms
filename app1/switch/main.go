package main

import "fmt"

func main() {

	var a, b, choice int
	fmt.Println("Enter 2 Numbers:")
	fmt.Scan(&a, &b)

	for {
		fmt.Println("1.Add\n2. Sub\n3.Mul\n0.Exit\nEnter Your Choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println(a + b)
		case 2:
			fmt.Println(a - b)
		case 3:
			fmt.Println(a * b)
		case 0:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
