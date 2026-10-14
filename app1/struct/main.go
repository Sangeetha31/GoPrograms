package main

import "fmt"

type Product struct {
	id       int
	name     string
	price    float64
	category string
}

func (prod Product) PrintProductDetails() {
	fmt.Println("Id:", prod.id, "Name:", prod.name, "Price:", prod.price, "Category:", prod.category)

}

func main() {
	prod1 := Product{
		id: 12101, name: "Iphone 15", price: 5600, category: "Mobiles",
	}

	prod1.PrintProductDetails()

	prod2 := Product{
		id: 12102, name: "Iphone 16", price: 5630, category: "Mobiles",
	}

	prod2.PrintProductDetails()

}
