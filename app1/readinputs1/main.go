package main

import "fmt"

func main() {

	var firstname, lastname string

	fmt.Println("Enter Your First Name:")
	fmt.Scanln(&firstname)

	fmt.Println("Enter Your Last Name:")
	fmt.Scanln(&lastname)

	fmt.Println("Welcone,", firstname, " ", lastname)
}
