package handlers

import (
	"net/http"
	"strconv"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"github.com/gin-gonic/gin"
)

// estrutura
type UsuarioHandle struct {
	services *services.UsuarioService
}

// construtor
func NewUsuarioHandle(s *services.UsuarioService) *UsuarioHandle {
	return &UsuarioHandle{
		services: s,
	}
}

// listar os usuários
func (u *UsuarioHandle) UsuarioListHandle(c *gin.Context) {
	// recebendo os valores via json
	users := u.services.UsuarioServiceList()

	// json
	c.JSON(http.StatusOK, gin.H{
		"message": "rota listar todos os usuario listar",
		"dados":   users,
	})
}

// inserir usuarios
func (u *UsuarioHandle) UsuarioPostHandle(c *gin.Context) {
	// recebendo os valores via json
	var user models.UsuarioCriate

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": err.Error(),
		})
		return
	}

	if saida := u.services.UsuarioServiceAdd(user); saida {
		// json
		c.JSON(http.StatusCreated, gin.H{
			"message": "rota registrar usuario",
			"dados":   saida,
		})
		return
	}

}

// byid usuarios
func (u *UsuarioHandle) UsuarioByIdHandle(c *gin.Context) {
	// recebendo os valores via json
	id, _ := strconv.Atoi(c.Param("id"))

	if user := u.services.UsuarioServiceById(int64(id)); user != nil {
		// json
		c.JSON(http.StatusOK, gin.H{
			"message": "rota um usuario",
			"dados":   user,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "usuario nao encontrado",
	})
	return
}

// update usuarios
func (u *UsuarioHandle) UsuarioUpdateHandle(c *gin.Context) {
	// recebendo os valores via json
	var user models.UsuarioCriate
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": err.Error(),
		})
		return
	}

	if saida := u.services.UsuarioServiceUpdate(id, user); saida {
		// json
		c.JSON(http.StatusOK, gin.H{
			"message": "rota update usuario",
			"dados":   user,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Usuario nao encontrado",
	})
}

// delete usuarios
func (u *UsuarioHandle) UsuarioDeleteHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if saida := u.services.UsuarioServiceDelete(id); saida {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	// json
	c.JSON(http.StatusBadRequest, gin.H{
		"mensagem": "usuario nao encontrado",
	})
}
