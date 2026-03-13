package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"github.com/gin-gonic/gin"
)

// estrutura que recebe os handlers
type RotasDefault struct {
	handle *handlers.DefaultHandle
}

// construtor
func NewRotasDefault(h *handlers.DefaultHandle) *RotasDefault {
	return &RotasDefault{handle: h}
}

// rota inicial
func (rou *RotasDefault) InicialRota(r *gin.Engine) {

	// rotas padrão
	r.GET("/", rou.handle.InicialHandle)
	r.NoRoute(rou.handle.NaoEncontrada)

	// agrupamento base
	r.Group("/")
	{
		// criando a base de usuarios
		repo := repository.NewUsuarioRepository()
		s := services.NewUsuarioService(repo)
		h := handlers.NewUsuarioHandle(s)
		user := NewUsuarioRotas(h)

		// rotas de usuários
		user.UsuarioRotas(r)
	}
}
