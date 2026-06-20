package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/BrunoPessoa097/api-go-simples/internal/start"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// estrutura que recebe os handlers
type RotasDefault struct {
	handle *handlers.DefaultHandle
	db     *gorm.DB
}

// construtor
func NewRotasDefault(h *handlers.DefaultHandle, db *gorm.DB) *RotasDefault {
	return &RotasDefault{
		handle: h,
		db:     db,
	}
}

// rota inicial
func (rou *RotasDefault) InicialRota(r *gin.Engine) {

	// rotas padrão
	r.GET("/", rou.handle.InicialHandle)
	r.NoRoute(rou.handle.NaoEncontrada)

	// agrupamento base
	r.Group("/")
	{
		//...
		// criando a base de usuarios
		st := start.NewStart(rou.db)
		uh, _ := st.UsuarioStart()
		user := NewUsuarioRotas(uh)

		// rotas de usuários
		user.UsuarioRotas(r)

		//rota roles
		rh := st.RoleStart()
		roles := NewRolesRoute(rh)
		roles.RolesRoutes(r)

		//post
		ps := st.PostStart()
		pp := NewPostRoutas(ps)
		pp.PostRotas(r)
	}
}
