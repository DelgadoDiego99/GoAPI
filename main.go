package main

import (
	"fmt"
	"todo/api/todo"

	"github.com/gin-gonic/gin"
)

func sayHi(c *gin.Context) {
	fmt.Println("Hello, world")
}

func main() {
	router := gin.Default()
	router.GET("/", sayHi)
	todoRoute := router.Group("/todo")
	{
		todoRoute.GET("/", todo.GetAllTodo)
	}
	router.Run("localhost:8080")
}