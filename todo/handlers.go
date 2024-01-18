package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTodo(c *gin.Context){
	c.IndentedJSON(http.StatusOK, lista)
}