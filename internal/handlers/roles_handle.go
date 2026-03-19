package handlers

import (
	"net/http"

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
	c.JSON(http.StatusCreated, gin.H{
		"mensage": "adicionar regras",
	})
}

// buscar por id
func (r *RolesHandler) RolesHandlerById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "buscar por id",
	})
}

// update
func (r *RolesHandler) RolesHandlerUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "update role",
	})
}

func (r *RolesHandler) RolesHandlerDelete(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
