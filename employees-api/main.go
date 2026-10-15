package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	Designation string  `json:"designation"`
	Salary      float32 `json:"salary"`
}

func getAllData(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			5*time.Second,
		)
		defer cancel()

		query := `SELECT * from employees`

		rows, err := db.QueryContext(ctx, query)
		if err != nil {
			log.Printf("Error: while retreiving data: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get data",
			})
			return
		}
		defer rows.Close()

		emps := make([]Employee, 0)

		for rows.Next() {
			var employee Employee

			err := rows.Scan(
				&employee.ID,
				&employee.Name,
				&employee.Age,
				&employee.Designation,
				&employee.Salary,
			)
			if err != nil {
				log.Printf("Error: scanning employee: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "failed to retrieve employees",
				})
				return
			}
			emps = append(emps, employee)
		}
		if err := rows.Err(); err != nil {
			log.Printf("Error: Reading employee rows: %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to Get Data",
			})
			return
		}
		c.JSON(http.StatusOK, emps)
	}
}

func addEmployee(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var employee Employee

		if err := ctx.ShouldBindJSON(&employee); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Request Body",
			})
			return
		}
		if employee.Name == "" || employee.Designation == "" || employee.Age <= 0 || employee.Salary <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Data",
			})
			return
		}
		cont, cancel := context.WithTimeout(
			ctx.Request.Context(), 5*time.Second,
		)

		defer cancel()

		query := `
		INSERT INTO employees(name,age,designation,salary) VALUES(?,?,?,?)
		`
		result, err := db.ExecContext(
			cont, query,
			employee.Name, employee.Age, employee.Designation, employee.Salary)
		if err != nil {
			log.Printf("Error: Inserting Data Into Employees: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to Insert Employee....",
			})
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Printf("Error: Getting Generated Employee ID: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to Insert Employee....",
			})
			return
		}
		employee.ID = int(id)

		ctx.JSON(http.StatusCreated, employee)
	}
}

func main() {

	db, err := connectDB()

	if err != nil {
		log.Fatal("ERROR: Connecting to DB: %v", err)
	}

	defer db.Close()

	router := gin.Default()

	router.GET("/employees", getAllData(db))

	router.POST("/employees", addEmployee(db))

	log.Println("Server is Up and Running on: 8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error: Starting the Server: %v", err)
	}
}
