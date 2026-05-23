package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandlers struct{}

func NewPostHandlers() *PostHandlers {
	return &PostHandlers{}
}

// listagem de postagem
func (p *PostHandlers) PostHandlersList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "listagem de postagem",
	})
}

// criar postagem
func (p *PostHandlers) PostHandlersPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Inserir de postagem",
	})
}

// selecionar postagem
func (p *PostHandlers) PostHandlersById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Buscar por postagem",
	})
}

// atualizar postagem
func (p *PostHandlers) PostHandlersUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "update de post",
	})
}

// deletar postagem
func (p *PostHandlers) PostHandlersDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "delete de post",
	})
}
