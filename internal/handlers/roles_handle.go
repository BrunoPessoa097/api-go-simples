package handlers

import (
	"net/http"
	"strconv"

	"github.com/BrunoPessoa097/api-go-simples/internal/dto"
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
	// validação
	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"mensage": "Listar regras",
			"dados":   "sem regras registradas",
		})
	}

	datas := dto.ToResponseRolesList(data)

	c.JSON(http.StatusOK, gin.H{
		"mensage": "Listar regras",
		"dados":   datas,
	})
}

// post
func (r *RolesHandler) RolesHandlerPost(c *gin.Context) {
	var input dto.RolesCreateDTO
	pkg := pkg.NewPkg()

	// validação de erros
	if err := c.BindJSON(&input); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erros,
		})
		return
	}

	roles := dto.ToModelRoles(input)

	//verificando a existencia
	if err := r.service.RoleServiceSearch(roles.Nivel); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"erro": err.Error(),
		})
		return
	}

	//cadastrando
	if err := r.service.RoleServicePost(roles); err != nil {
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
	if saida, _ := r.service.RoleServiceById(idC); saida != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Regras encontrado",
			"dado":    dto.ToResponseRoles(*saida),
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

	exist, _ := r.service.RoleServiceById(idC)
	if exist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Regras nao encontrado",
		})
		return
	}

	var input dto.RolesUpdateDTO
	// validação de erros
	if err := c.BindJSON(&input); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erros,
		})
		return
	}

	//verificando a existencia
	if input.Nivel != nil {
		if err := r.service.RoleServiceSearch(*input.Nivel); err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"erro": err.Error(),
			})
			return
		}
	}

	update := dto.ToUpdateRoles(input, *exist)

	if saida := r.service.RoleServiceUpdate(idC, update); saida != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": saida.Error(),
		})
		return
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
	if _, erro := r.service.RoleServiceById(idC); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Regras nao encontrado",
		})
		return
	}

	//saida
	if err := r.service.RoleServiceDelete(idC); err != nil {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	//saida sem sucesso
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Regras não encontrado",
	})
}
