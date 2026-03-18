package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RolesHandler struct{}

func NewRolesHandler() *RolesHandler {
	return &RolesHandler{}
}

func (r *RolesHandler) RolesHandlerList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensage": "Listar regras",
	})
}
