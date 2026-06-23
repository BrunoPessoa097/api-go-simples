package start

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
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
func (s *Start) UsuarioStart() (*handlers.UsuarioHandle, *repository.RolesRepository) {
	repo := repository.NewUsuarioRepository(s.db)
	rr := repository.NewRolesRepository(s.db)
	serv := services.NewUsuarioService(repo, rr)
	hand := handlers.NewUsuarioHandle(serv)
	return hand, rr
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
	pr := repository.NewPostRepository(s.db)
	ru := repository.NewUsuarioRepository(s.db)
	ps := services.NewPostService(pr, ru)
	ph := handlers.NewPostHandlers(ps)
	return ph
}

// inicio do login
func (s *Start) LoginStart() *handlers.Importante {
	ru := repository.NewUsuarioRepository(s.db)
	return handlers.NewImporteHandle(ru)
}
