package main

import "fmt"

type Address struct {
	city  string
	state string
	pin   string
}

type Employee struct {
	id          int
	name        string
	age         int
	designation string
	address     Address
}

func main() {
	emp := Employee{
		id: 12580, age: 25, name: "Ram", designation: "Manager",
		address: Address{
			city: "Mumbai", state: "MH", pin: "250255",
		},
	}
	fmt.Println(emp.id, emp.name, emp.address.city)
}
