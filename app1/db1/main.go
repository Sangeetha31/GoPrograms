package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := os.Getenv("MYSQL_DSN")

	if dsn == "" {
		fmt.Println("MYSQL DSN is not SET")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error While Opening Connection:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return
	}

	fmt.Println("Connected to DB!")
}
