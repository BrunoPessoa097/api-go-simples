package start

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"gorm.io/gorm"
)

// estrutura
type Start struct {
	db *gorm.DB
}

// construtor
func NewStart(db *gorm.DB) *Start {
	return &Start{
		db: db,
	}
}

// inicializando o usuarios
func (s *Start) UsuarioStart() *handlers.UsuarioHandle {
	repo := repository.NewUsuarioRepository(s.db)
	serv := services.NewUsuarioService(repo)
	hand := handlers.NewUsuarioHandle(serv)
	return hand
}

// iniciando regras
func (s *Start) RoleStart() *handlers.RolesHandler {
	rr := repository.NewRolesRepository(s.db)
	rs := services.NewRoleService(rr)
	rh := handlers.NewRolesHandler(rs)

	return rh
}

// iniciando postagem
func (s *Start) PostStart() *handlers.PostHandlers {
	mocks := mocks.ListPost
	pr := repository.NewPostRepository(mocks)
	ps := services.NewPostService(pr)
	ph := handlers.NewPostHandlers(ps)
	return ph
}
