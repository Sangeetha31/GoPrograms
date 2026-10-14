package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDBConnection() (*sql.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/storedb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	fmt.Println("Connected to DB!")
	return db, err
}
