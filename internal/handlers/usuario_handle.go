package handlers

import (
	"net/http"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/gin-gonic/gin"
)

func UsuarioListHandle(c *gin.Context) {
	var user models.UsuarioCriate
	c.JSON(http.StatusOK, gin.H{
		"message": "rota usuario listar",
		"dados":   user,
	})
}
