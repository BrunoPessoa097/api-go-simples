package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var dtHr = time.Now().Format("01-02-2006 15:04:05")

func InicialHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem":  "API simples em golang(estudos 3/3)",
		"versao":    "v0.0.0",
		"autor":     "Bruno Pessoa",
		"data/hora": dtHr,
	})
}

func NaoEncontrada(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"menssage":  "rota não encontrada",
		"data/hora": dtHr,
	})
}
