package main

import (
	"context"
	"database/sql"
	"time"
)

func connectDB() (*sql.DB, error) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/officedb"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second,
	)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(3 * time.Minute)

	return db, nil
}
