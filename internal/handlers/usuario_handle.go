package handlers

import (
	"net/http"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/gin-gonic/gin"
)

// estrutura
type UsuarioHandle struct{}

// construtor
func NewUsuarioHandle() *UsuarioHandle {
	return &UsuarioHandle{}
}

// listar os usuários
func (u *UsuarioHandle) UsuarioListHandle(c *gin.Context) {
	// recebendo os valores via json
	var user models.UsuarioCriate

	// json
	c.JSON(http.StatusOK, gin.H{
		"message": "rota listar todos os usuario listar",
		"dados":   user,
	})
}

// inserir usuarios
func (u *UsuarioHandle) UsuarioPostHandle(c *gin.Context) {
	// recebendo os valores via json
	var user models.UsuarioCriate

	// json
	c.JSON(http.StatusOK, gin.H{
		"message": "rota registrar usuario",
		"dados":   user,
	})
}

// inserir usuarios
func (u *UsuarioHandle) UsuarioByIdHandle(c *gin.Context) {
	// recebendo os valores via json
	var user models.UsuarioCriate

	// json
	c.JSON(http.StatusOK, gin.H{
		"message": "rota um usuario",
		"dados":   user,
	})
}

// inserir usuarios
func (u *UsuarioHandle) UsuarioUpdateHandle(c *gin.Context) {
	// recebendo os valores via json
	var user models.UsuarioCriate

	// json
	c.JSON(http.StatusOK, gin.H{
		"message": "rota update usuario",
		"dados":   user,
	})
}
