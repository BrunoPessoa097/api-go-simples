package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InicialHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "API simples em golang(estudos)",
		"versao":   "v0.0.0",
		"autor":    "Bruno Pessoa",
	})
}
