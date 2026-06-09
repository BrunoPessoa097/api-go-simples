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
	datas := p.Service.PostServiceList()
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

	if ok := p.Service.PostServiceAdd(*post); ok {
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "Inserir de postagem",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"mensagem": "Post nao inserido",
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

	update := dto.ToModelPostUpdade(input)

	if saida := p.Service.PostServiceUpdate(idCvt, update); saida != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "update de post",
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
	if saida := p.Service.PostServiceDelete(idCvt); saida != nil {
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "Postagem não encontrada",
			"erro":     saida.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "delete de post com sucesso",
	})
}
