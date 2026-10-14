package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := createTable(); err != nil {
		fmt.Println("Error:", err)
	}

	/*if err := Insert(); err != nil {
		fmt.Println("Error:", err)
	}*/

	if err := getData(); err != nil {
		fmt.Println("Error:", err)
	}
}

func createTable() error {
	db, err := getDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
		CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		price FLOAT NOT NULL
		)
		`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Table Created Successfully")
	return nil
}

func Insert() error {
	db, err := getDBConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Product Name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	name = strings.TrimSpace(name)

	if name == "" {
		return fmt.Errorf("Name Should not be Empty!")
	}

	fmt.Println("Enter Product Price:")
	priceInp, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	price, err := strconv.ParseFloat(strings.TrimSpace(priceInp), 64)
	if err != nil || price < 0 {
		return fmt.Errorf("Invalid Price")
	}

	query := `
	INSERT INTO products(name,price) VALUES(?,?)
	`

	res, err := db.Exec(query, name, price)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println("\nProduct Data Saved Successfully!")
	fmt.Println(id, name, price)

	return nil
}

func getData() error {
	db, err := getDBConnection()
	if err != nil {
		return err
	}

	defer db.Close()

	query :=
		`SELECT * FROM products`

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	defer rows.Close()

	fmt.Printf("%-5s %-20s %-12s\n", "ID", "Name", "Price")
	fmt.Println(strings.Repeat("-", 50))

	rowcount := 0

	for rows.Next() {
		var id int
		var name string
		var price float64

		err := rows.Scan(&id, &name, &price)
		if err != nil {
			return err
		}
		rowcount++

		fmt.Printf("%-5d %-20s %-12.0f", id, name, price)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if rowcount == 0 {
		fmt.Println("No Data Found!")
	}

	return nil
}
