package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Produto struct {
	Nome  string  `json:"nome" binding:"required"`
	Preco float64 `json:"preco" binding:"required,gt=0"`
}

func main() {
	r := gin.Default()
	r.POST("/produtos", func(c *gin.Context) {
		var p Produto
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	})

	r.Run(":8080")
}
