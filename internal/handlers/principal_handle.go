package handlers

import (
	"net/http"

	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/gin-gonic/gin"
)

// estrutura principal
type DefaultHandle struct{}

// construtor
func NewDefaultHandle() *DefaultHandle {
	return &DefaultHandle{}
}

// saida da rota principal
func (d *DefaultHandle) InicialHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem":  "API simples em golang(estudos 3/3)",
		"versao":    "v0.0.0",
		"autor":     "Bruno Pessoa",
		"data/hora": utils.DtHr,
	})
}

// saida da rota não encontrada
func (d *DefaultHandle) NaoEncontrada(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"menssage":  "rota não encontrada",
		"data/hora": utils.DtHr,
	})
}
