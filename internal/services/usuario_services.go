package services

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
)

type UsuarioService struct {
	repo *repository.UsuarioRepository
}

// construtor
func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		repo: repo,
	}
}

// listar
func (s *UsuarioService) UsuarioServiceList() ([]models.Usuario, error) {
	return s.repo.UsuarioRepositoryList()
}

// adiciona
func (s *UsuarioService) UsuarioServiceAdd(user models.Usuario) error {
	return s.repo.UsuarioRepositoryAdd(user)
}

// buscar por id
func (s *UsuarioService) UsuarioServiceById(id int32) (*models.Usuario, error) {
	return s.repo.UsuarioRepositoryById(id)
}

// update
func (s *UsuarioService) UsuarioServiceUpdate(id int32, user models.Usuario) error {
	user.ID = int32(id)
	return s.repo.UsuarioRepositoryUpdate(user)
}

// delete
func (s *UsuarioService) UsuarioServiceDelete(id int32) error {
	return s.repo.UsuarioRepositoryDelete(id)
}
