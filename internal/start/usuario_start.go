package start

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
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
	repo := repository.NewUsuarioRepository()
	serv := services.NewUsuarioService(repo)
	hand := handlers.NewUsuarioHandle(serv)
	return hand
}
