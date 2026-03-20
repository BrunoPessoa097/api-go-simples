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

	// enviando dados
	_, err := r.service.RoleServicePost(roles)

	// saida de erros
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensage": err.Error(),
		})
		return
	}

	//saida de sucesso
	c.JSON(http.StatusCreated, gin.H{
		"mensage": "regra adicionada",
		"dados":   roles,
	})
}

// buscar por id
func (r *RolesHandler) RolesHandlerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	dado, err := r.service.RoleServiceById(int64(id))

	if dado == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "buscar por id",
		"dado":    dado,
	})
}

// update
func (r *RolesHandler) RolesHandlerUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "update role",
	})
}

// delete
func (r *RolesHandler) RolesHandlerDelete(c *gin.Context) {
	//conversoes
	id, _ := strconv.Atoi(c.Param("id"))
	ok := r.service.RoleServiceDelete(int64(id))

	//em caso de erro
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "erro ao excluir",
		})
		return
	}
	//saida um sucesso
	c.JSON(http.StatusNoContent, nil)
}
