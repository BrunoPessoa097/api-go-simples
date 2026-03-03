package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// data so servidor
var dtHr = time.Now().Format("01-02-2006 15:04:05")

// saida da rota principal
func InicialHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem":  "API simples em golang(estudos 3/3)",
		"versao":    "v0.0.0",
		"autor":     "Bruno Pessoa",
		"data/hora": dtHr,
	})
}

// saida da rota não encontrada
func NaoEncontrada(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"menssage":  "rota não encontrada",
		"data/hora": dtHr,
	})
}
