package main

import (
	"fmt"
	"net/http"
	"todo/api/todo"

	"github.com/gin-gonic/gin"
)

func sayHi(c *gin.Context) {
	fmt.Println("Hello, world")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("*.html")

	router.GET("/", func(c *gin.Context) {c.HTML(http.StatusOK, "index.html", "")})
	todoRoute := router.Group("/todo")
	{
		todoRoute.GET("/", todo.GetAllTodo)
	}
	router.Run("localhost:8080")
}