package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Funcao para rota Default Padrão
func inicial(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Rota inicial",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", inicial)

	r.Run(":8080")
}
