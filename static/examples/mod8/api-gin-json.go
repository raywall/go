package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.Header("X-Custom-Header", "valor")
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "API funcionando",
		})
	})

	r.Run(":8080")
}
