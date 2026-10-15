package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:id`
	Name string `json:name`
}

func main() {
	router := gin.Default()

	router.GET("/users", func(ctx *gin.Context) {
		users := []User{
			{ID: 1258, Name: "Veera"},
			{ID: 16859, Name: "Harry"},
		}
		ctx.JSON(http.StatusOK, users)
	})

	router.POST("/users", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Data",
			})
		}
		ctx.JSON(http.StatusCreated, user)
	})

	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}
