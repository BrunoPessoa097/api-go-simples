package repository

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

type UsuarioRepository struct{}

// construtor
func NewUsuarioRepository() *UsuarioRepository {
	return &UsuarioRepository{}
}

// listar os usuarios
func (r *UsuarioRepository) UsuarioRepositoryList() []models.UsuarioCriate {
	return mocks.UsuariosBD
}
