package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	products := []Product{
		{ID: 1, Name: "Product-1", Price: 5000},
		{ID: 2, Name: "Product-2", Price: 7500},
		{ID: 3, Name: "Product-3", Price: 6600},
		{ID: 4, Name: "Product-4", Price: 7700},
		{ID: 5, Name: "Product-5", Price: 3200},
	}

	jsondata, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		fmt.Println("Error While Converting to JSON:", err)
		return
	}

	err = os.WriteFile("products.json", jsondata, 0644)
	if err != nil {
		fmt.Println("Error While Writing JSON File:", err)
		return
	}
	fmt.Println("Data Saved to JSON File")
}
