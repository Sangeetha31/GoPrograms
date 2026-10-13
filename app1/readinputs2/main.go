package main

import "fmt"

func main() {

	var firstname, lastname string

	fmt.Println("Enter Your First Name:")
	fmt.Scanf("%[^\n]", &firstname)

	fmt.Println("Enter Your Last Name:")
	fmt.Scanf("%[^[\n]", &lastname)

	fmt.Println("Welcone,", firstname, " ", lastname)
}
