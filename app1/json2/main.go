package main

import (
	"encoding/json/v2"
	"fmt"
	"os"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	data, err := os.ReadFile("products.json")

	if err != nil {
		fmt.Println("Error While Reading File:", err)
		return
	}

	var products []Product

	json.Unmarshal(data, &products)
	if err != nil {
		fmt.Println("Error While Parsing JSON", err)
		return
	}

	for _, prod := range products {
		fmt.Printf(
			"ID: %d,Name:%s,Price: %.2f\n",
			prod.ID,
			prod.Name,
			prod.Price,
		)
	}
}
