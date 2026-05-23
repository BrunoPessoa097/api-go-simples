package start

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
)

// estrutura
type Start struct {
}

// construtor
func NewStart() *Start {
	return &Start{}
}

// inicializando o usuarios
func (s *Start) UsuarioStart() *handlers.UsuarioHandle {
	mock := mocks.UsuariosBD
	repo := repository.NewUsuarioRepository(mock)
	serv := services.NewUsuarioService(repo)
	hand := handlers.NewUsuarioHandle(serv)
	return hand
}

// iniciando regras
func (s *Start) RoleStart() *handlers.RolesHandler {
	mocks := mocks.ListRoles
	rr := repository.NewRolesRepository(mocks)
	rs := services.NewRoleService(rr)
	rh := handlers.NewRolesHandler(rs)

	return rh
}

// iniciando postagem
func (s *Start) PostStart() *handlers.PostHandlers {
	ph := handlers.NewPostHandlers()
	return ph
}
