package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/gin-gonic/gin"
)

// estrutura
type UsuarioRotes struct {
	handle *handlers.UsuarioHandle
}

// construtor
func NewUsuarioRotas(h *handlers.UsuarioHandle) *UsuarioRotes {
	return &UsuarioRotes{
		handle: h,
	}
}

// roteamento de usuários
func (u *UsuarioRotes) UsuarioRotas(rg *gin.Engine) {

	// grupo das rotas
	usuario := rg.Group("/usuario")
	{
		usuario.GET("/", u.handle.UsuarioListHandle)
		usuario.POST("/", u.handle.UsuarioPostHandle)
	}
}
