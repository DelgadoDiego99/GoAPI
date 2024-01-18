package main

import (
	"todo/api/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", todo.GetAllTodo)
	router.Run("localhost:8080")
}