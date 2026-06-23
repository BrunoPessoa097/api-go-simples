package handlers

import (
	"net/http"

	"github.com/BrunoPessoa097/api-go-simples/internal/dto"
	"github.com/BrunoPessoa097/api-go-simples/internal/pkg"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/gin-gonic/gin"
)

type Importante struct {
	repo *repository.UsuarioRepository
}

// construtor
func NewImporteHandle(ser *repository.UsuarioRepository) *Importante {
	return &Importante{
		repo: ser,
	}
}

// login
func (i *Importante) ImportanteLogin(c *gin.Context) {
	var login dto.LoginCreate
	pkg := pkg.NewPkg()

	if err := c.BindJSON(&login); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"dados": erros,
		})
		return
	}

	loginDTO := dto.ToLogin(&login)

	saida, err := i.repo.UsuarioRepositorySearch("", loginDTO.Email)

	if !err {
		c.JSON(http.StatusOK, gin.H{
			"menssage": "e-mail não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menssage": "login",
		"dados":    saida,
	})
}
