package handlers

import (
	"net/http"
	"strconv"

	"github.com/BrunoPessoa097/api-go-simples/internal/dto"
	"github.com/BrunoPessoa097/api-go-simples/internal/pkg"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
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
	user, _ := u.services.UsuarioServiceList()

	if len(user) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "rota listar todos os usuario listar",
			"dados":   "Sem usuários inseridos",
		})
		return
	}

	users := dto.ToResponseList(user)

	// json
	c.JSON(http.StatusOK, gin.H{
		"message": "rota listar todos os usuario listar",
		"dados":   users,
	})
}

// inserir usuarios
func (u *UsuarioHandle) UsuarioPostHandle(c *gin.Context) {
	// recebendo os valores via json
	var input dto.UsuarioCriateDTO
	pkg := pkg.NewPkg()

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := pkg.Validator(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": errors,
		})
		return
	}

	user := dto.ToModel(input)

	user.Senha = utils.Hash(user.Senha)

	if err := u.services.UsuarioServiceAdd(user); err != nil {
		// json
		c.JSON(http.StatusConflict, gin.H{
			"message": "erro ao cadastrar usuário",
			"dados":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "usuario cadastrado",
	})

}

// byid usuarios
func (u *UsuarioHandle) UsuarioByIdHandle(c *gin.Context) {
	// recebendo os valores via json
	id, _ := strconv.Atoi(c.Param("id"))

	if user, _ := u.services.UsuarioServiceById(int32(id)); user != nil {
		saida := dto.ToResponse(*user)
		// json
		c.JSON(http.StatusOK, gin.H{
			"message": "rota um usuario",
			"dados":   saida,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "usuario nao encontrado",
	})
}

// update usuarios
func (u *UsuarioHandle) UsuarioUpdateHandle(c *gin.Context) {
	// recebendo os valores via json
	id, _ := strconv.Atoi(c.Param("id"))

	if err, _ := u.services.UsuarioServiceById(int32(id)); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Usuario update",
			"dados":   "usuario não encontrado",
		})
		return
	}

	var user dto.UsuarioUpdateDTO
	pkg := pkg.NewPkg()

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		erros := pkg.Validator(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": erros,
		})
		return
	}

	users := dto.ToUpdate(user)

	if err := u.services.UsuarioServiceUpdate(int32(id), users); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"menssage": "erro update",
			"erro":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menssage": "Update usuario",
	})
}

// delete usuarios
func (u *UsuarioHandle) UsuarioDeleteHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err, _ := u.services.UsuarioServiceById(int32(id)); err == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Mensasse": "usuario não encontrado",
		})
		return
	}

	if saida := u.services.UsuarioServiceDelete(int32(id)); saida != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "erro ao excluir usuario",
			"erro":     saida.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
