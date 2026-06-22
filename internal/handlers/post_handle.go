package handlers

import (
	"net/http"
	"strconv"

	"github.com/BrunoPessoa097/api-go-simples/internal/dto"
	"github.com/BrunoPessoa097/api-go-simples/internal/pkg"
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
	datas, _ := p.Service.PostServiceList()

	if len(datas) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"mensage": "Listar regras",
			"dados":   "sem regras registradas",
		})
		return
	}

	data := dto.ToModelPostList(datas)
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "listagem de postagem",
		"data":     data,
	})
}

// criar postagem
func (p *PostHandlers) PostHandlersPost(c *gin.Context) {
	var input dto.CreatePostDTO
	pkg := pkg.NewPkg()

	if err := c.BindJSON(&input); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erros,
		})
		return
	}

	post := dto.ToModelPostCreate(input)

	if ok := p.Service.PostServiceAdd(post); ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "Post nao inserido",
			"erro":     ok.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"mensagem": "Post inserido",
	})
}

// selecionar postagem
func (p *PostHandlers) PostHandlersById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := p.Service.PostRepositoryId(int64(id))

	if post != nil {
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "Buscar por postagem",
			"dado":     dto.ToModelPostListOne(*post),
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"mensagem": "Buscar por postagem err",
		"erro":     err.Error(),
	})
}

// atualizar postagem
func (p *PostHandlers) PostHandlersUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idCvt := int64(id)
	pkg := pkg.NewPkg()

	var input dto.UpdatePostDTO

	if err := c.BindJSON(&input); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erros,
		})
		return
	}

	post, _ := p.Service.PostRepositoryId(idCvt)

	ui := int64(post.IDUser)
	input.IDUser = &ui

	update := dto.ToModelPostUpdade(input)

	if saida := p.Service.PostServiceUpdate(idCvt, update); saida != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "update de post erro",
			"erros":    saida.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "update na postagem",
	})
}

// deletar postagem
func (p *PostHandlers) PostHandlersDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idCvt := int64(id)

	if err := p.Service.PostServiceDelete(idCvt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Postagem não encontrada",
			"erro":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "delete de post com sucesso",
	})
}
