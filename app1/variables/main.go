package main

import "fmt"

var appName string = "MyAPP"
var appVer float32 = 1.0

func main() {
	var a int = 10

	var b int = 20

	b = b + 1
	fmt.Println("The value of a is:", a)
	fmt.Println("The value of b is:", b)

	fmt.Printf("The value of a is: %d", a)
	fmt.Printf("The value of b is: %d", b)

	var score float32
	fmt.Println("The Value in the score", score)

	var isLoggedInn bool
	fmt.Println("Login Status:", isLoggedInn)

	var nationality string
	fmt.Println("Nationality", nationality)

	name := "Veer"
	age := 28
	city := "Yelahanka, Bengaluru"

	fmt.Println("Name:", name, "\tAge:", age, "\tCity:", city)

	designation, location := "Product-Specialist", "Bengaluru"
	id, salary := 25896, 125000

	fmt.Println("Id:", id, "\tname:", name, "\tDesignation:", designation, "\tBase location: ", location, "\tSalary:", salary)

	var (
		EmpCode int    = 2589
		EmpName string = "Veer"
	)

	fmt.Println("Code:", EmpCode, "\tName:", EmpName)

	fmt.Println("APP Deatils:\t Name:", appName, "\t Version:", appVer)
}
