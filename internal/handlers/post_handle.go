package handlers

import (
	"net/http"

	"github.com/BrunoPessoa097/api-go-simples/internal/dto"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"github.com/gin-gonic/gin"
)

type PostHandlers struct {
	Service *services.PostService
}

func NewPostHandlers(s *services.PostService) *PostHandlers {
	return &PostHandlers{
		Service: s,
	}
}

// listagem de postagem
func (p *PostHandlers) PostHandlersList(c *gin.Context) {
	datas := p.Service.PostServiceList()
	data := dto.ToModelPostList(datas)
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "listagem de postagem",
		"data":     data,
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
