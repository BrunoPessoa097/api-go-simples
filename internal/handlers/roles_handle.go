package handlers

import (
	"net/http"
	"strconv"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/pkg"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"github.com/gin-gonic/gin"
)

// estrutura
type RolesHandler struct {
	service *services.RoleService
}

// construtor
func NewRolesHandler(s *services.RoleService) *RolesHandler {
	return &RolesHandler{service: s}
}

// list
func (r *RolesHandler) RolesHandlerList(c *gin.Context) {
	data := r.service.RoleServiceList()
	c.JSON(http.StatusOK, gin.H{
		"mensage": "Listar regras",
		"dados":   data,
	})
}

// post
func (r *RolesHandler) RolesHandlerPost(c *gin.Context) {
	var roles models.RolesC
	pkg := pkg.NewPkg()

	// validação de erros
	if err := c.BindJSON(&roles); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erros,
		})
		return
	}

	//verificando a existencia
	if err := r.service.RoleServiceSearch(roles.Nivel); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"erro": err.Error(),
		})
		return
	}

	//cadastrando
	if err := r.service.RoleServicePost(roles); !err {
		c.JSON(http.StatusBadRequest, gin.H{
			"menssagem": "erro ao cadastrar",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"menssagem": "Regra cadastrado",
	})
}

// buscar por id
func (r *RolesHandler) RolesHandlerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idC := int64(id)

	// saida
	if saida := r.service.RoleServiceById(idC); saida != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Regras encontrado",
			"dado":    saida,
		})
		return
	}

	// caso não achado
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "regra não encontrado",
	})
}

// // update
func (r *RolesHandler) RolesHandlerUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idC := int64(id)
	pkg := pkg.NewPkg()

	if saida := r.service.RoleServiceById(idC); saida == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Regras nao encontrado",
		})
		return
	}

	var roles models.RolesC
	// validação de erros
	if err := c.BindJSON(&roles); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erros,
		})
		return
	}

	//verificando a existencia
	if err := r.service.RoleServiceSearch(roles.Nivel); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"erro": err.Error(),
		})
		return
	}

	if saida := r.service.RoleServiceUpdate(idC, roles); saida != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": saida.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "update role",
	})
}

// // delete
func (r *RolesHandler) RolesHandlerDelete(c *gin.Context) {
	//conversoes
	id, _ := strconv.Atoi(c.Param("id"))
	idC := int64(id)

	//buscando id
	if saida := r.service.RoleServiceById(idC); saida == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Regras nao encontrado",
		})
		return
	}

	//saida
	if err := r.service.RoleServiceDelete(idC); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Regras nao",
		})
		return
	}

	//saida sem sucesso
	c.JSON(http.StatusNoContent, nil)
}
