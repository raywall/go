package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:nome", func(c *gin.Context) {
		nome := c.Param("nome")
		c.JSON(http.StatusOK, gin.H{"mensagem": "Olá, " + nome})
	})

	r.Run(":8080")
}
